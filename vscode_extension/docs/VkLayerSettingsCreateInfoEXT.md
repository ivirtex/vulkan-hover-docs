# VkLayerSettingsCreateInfoEXT(3) Manual Page

## Name

VkLayerSettingsCreateInfoEXT - Specify layer capabilities for a Vulkan instance



## [](#_c_specification)C Specification

To create a Vulkan instance with a specific configuration of layer settings, add [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html) structures to the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure, specifying the settings to be configured.

```c++
// Provided by VK_EXT_layer_settings
typedef struct VkLayerSettingsCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    uint32_t                    settingCount;
    const VkLayerSettingEXT*    pSettings;
} VkLayerSettingsCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `settingCount` is the number of settings to configure.
- `pSettings` is a pointer to an array of `settingCount` [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingEXT.html) values specifying the settings to be configured.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkLayerSettingsCreateInfoEXT-sType-sType)VUID-VkLayerSettingsCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LAYER_SETTINGS_CREATE_INFO_EXT`
- [](#VUID-VkLayerSettingsCreateInfoEXT-pSettings-parameter)VUID-VkLayerSettingsCreateInfoEXT-pSettings-parameter  
  If `settingCount` is not `0`, `pSettings` **must** be a valid pointer to an array of `settingCount` valid [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingEXT.html) structures

## [](#_see_also)See Also

[VK\_EXT\_layer\_settings](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_layer_settings.html), [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLayerSettingsCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0