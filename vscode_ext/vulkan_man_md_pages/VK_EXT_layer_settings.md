# VK_EXT_layer_settings(3) Manual Page

## Name

VK_EXT_layer_settings - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

497

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Christophe Riccio <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_layer_settings%5D%20@christophe%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_layer_settings%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>christophe</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_layer_settings](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_layer_settings.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-09-23

**IP Status**  
No known IP claims.

**Contributors**  
- Christophe Riccio, LunarG

- Mark Lobodzinski, LunarG

- Charles Giessen, LunarG

- Spencer Fricke, LunarG

- Juan Ramos, LunarG

- Daniel Rakos, RasterGrid

- Shahbaz Youssefi, Google

- Lina Versace, Google

- Bill Hollings, The Brenwill Workshop

- Jon Leech, Khronos

- Tom Olson, Arm

## <a href="#_description" class="anchor"></a>Description

This extension provides a mechanism for configuring programmatically
through the Vulkan API the behavior of layers.

This extension provides the
[VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html) struct
that can be included in the `pNext` chain of the
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure passed as
the `pCreateInfo` parameter of
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html).

The structure contains an array of
[VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html) structure values that
configure specific features of layers.

## <a href="#_example" class="anchor"></a>Example

`VK_EXT_layer_settings` is implemented by the Vulkan Profiles layer.

It allows the profiles layer tests used by the profiles layer C.I. to
programmatically configure the layer for each test without affecting the
C.I. environment, allowing to run multiple tests concurrently.

``` c
const char* profile_file_data = JSON_TEST_FILES_PATH "VP_KHR_roadmap_2022.json";
const char* profile_name_data = "VP_KHR_roadmap_2022";
VkBool32 emulate_portability_data = VK_TRUE;
const char* simulate_capabilities[] = {
    "SIMULATE_API_VERSION_BIT",
    "SIMULATE_FEATURES_BIT",
    "SIMULATE_PROPERTIES_BIT",
    "SIMULATE_EXTENSIONS_BIT",
    "SIMULATE_FORMATS_BIT",
    "SIMULATE_QUEUE_FAMILY_PROPERTIES_BIT"
};
const char* debug_reports[] = {
    "DEBUG_REPORT_ERROR_BIT",
    "DEBUG_REPORT_WARNING_BIT",
    "DEBUG_REPORT_NOTIFICATION_BIT",
    "DEBUG_REPORT_DEBUG_BIT"
};

const VkLayerSettingEXT settings[] = {
     {kLayerName, kLayerSettingsProfileFile, VK_LAYER_SETTING_TYPE_STRING_EXT, 1, &profile_file_data},
     {kLayerName, kLayerSettingsProfileName, VK_LAYER_SETTING_TYPE_STRING_EXT, 1, &profile_name_data},
     {kLayerName, kLayerSettingsEmulatePortability, VK_LAYER_SETTING_TYPE_BOOL32_EXT, 1, &emulate_portability_data},
     {kLayerName, kLayerSettingsSimulateCapabilities, VK_LAYER_SETTING_TYPE_STRING_EXT,
        static_cast<uint32_t>(std::size(simulate_capabilities)), simulate_capabilities},
     {kLayerName, kLayerSettingsDebugReports, VK_LAYER_SETTING_TYPE_STRING_EXT,
        static_cast<uint32_t>(std::size(debug_reports)), debug_reports}
};

const VkLayerSettingsCreateInfoEXT layer_settings_create_info{
    VK_STRUCTURE_TYPE_LAYER_SETTINGS_CREATE_INFO_EXT, nullptr,
    static_cast<uint32_t>(std::size(settings)), settings};

VkInstanceCreateInfo inst_create_info = {};
...
inst_create_info.pNext = &layer_settings_create_info;
vkCreateInstance(&inst_create_info, nullptr, &_instances);
```

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <a
href="VK_EXT_layer_settings.html"><code>VK_EXT_layer_settings</code></a>
extension subsumes all the functionality provided in the <a
href="VK_EXT_validation_flags.html"><code>VK_EXT_validation_flags</code></a>
extension and the <a
href="VK_EXT_validation_features.html"><code>VK_EXT_validation_features</code></a>
extension.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html)

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html):

  - [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingTypeEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_LAYER_SETTINGS_EXTENSION_NAME`

- `VK_EXT_LAYER_SETTINGS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_LAYER_SETTINGS_CREATE_INFO_EXT`

## <a href="#_issues" class="anchor"></a>Issues

- How should application developers figure out the list of available
  settings?

This extension does not provide a reflection API for layer settings.
Layer settings are described in each layer JSON manifest and the
documentation of each layer which implements this extension.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-06-17 (Mark Lobodzinski)

  - Initial revision for Validation layer internal usages

- Revision 2, 2023-09-26 (Christophe Riccio)

  - Refactor APIs for any layer usages and public release

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_layer_settings"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
