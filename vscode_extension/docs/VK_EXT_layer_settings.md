# VK\_EXT\_layer\_settings(3) Manual Page

## Name

VK\_EXT\_layer\_settings - instance extension



## [](#_registered_extension_number)Registered Extension Number

497

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Christophe Riccio [\[GitHub\]christophe](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_layer_settings%5D%20%40christophe%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_layer_settings%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_layer\_settings](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_layer_settings.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

This extension provides a mechanism for configuring programmatically through the Vulkan API the behavior of layers.

This extension provides the [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html) structure that can be included in the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure passed as the `pCreateInfo` parameter of [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html).

The structure contains an array of [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingEXT.html) structure values that configure specific features of layers.

Note

The `VK_EXT_layer_settings` extension subsumes all the functionality provided in the `VK_EXT_validation_flags` extension and the `VK_EXT_validation_features` extension.

## [](#_new_structures)New Structures

- [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingEXT.html)
- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html):
  
  - [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingTypeEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_LAYER_SETTINGS_EXTENSION_NAME`
- `VK_EXT_LAYER_SETTINGS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_LAYER_SETTINGS_CREATE_INFO_EXT`

## [](#_example)Example

One example usage of `VK_EXT_layer_settings` is as implemented by the Vulkan Profiles layer.

It allows the profiles layer tests used by the profiles layer C.I. to programmatically configure the layer for each test without affecting the C.I. environment, allowing to run multiple tests concurrently.

```c++
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

## [](#_issues)Issues

- How should application developers figure out the list of available settings?

This extension does not provide a reflection API for layer settings. Layer settings are described in each layer JSON manifest and the documentation of each layer which implements this extension.

## [](#_version_history)Version History

- Revision 1, 2020-06-17 (Mark Lobodzinski)
  
  - Initial revision for Validation layer internal usages
- Revision 2, 2023-09-26 (Christophe Riccio)
  
  - Refactor APIs for any layer usages and public release

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_layer_settings)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0