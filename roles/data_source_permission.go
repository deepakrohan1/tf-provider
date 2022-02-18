package roles

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePermissions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePermissionsRead,
		Schema: map[string]*schema.Schema{
			"permissions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"projectid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"projectname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"privilege": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePermissionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/roles", "http://localhost:8080"), nil)

	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}

	defer r.Body.Close()

	permissions := make([]map[string]interface{}, 0)

	err = json.NewDecoder(r.Body).Decode(&permissions)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("permissions", permissions); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags

}
