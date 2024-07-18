# VkLayerSettingsCreateInfoEXT(3) Manual Page

## Name

VkLayerSettingsCreateInfoEXT - Specify layer capabilities for a Vulkan
instance



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a Vulkan instance with a specific configuration of layer
settings, add
[VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html)
structures to the `pNext` chain of the
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure, specifying
the settings to be configured.

``` c
// Provided by VK_EXT_layer_settings
typedef struct VkLayerSettingsCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    uint32_t                    settingCount;
    const VkLayerSettingEXT*    pSettings;
} VkLayerSettingsCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `settingCount` is the number of settings to configure.

- `pSettings` is a pointer to an array of `settingCount`
  [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html) values specifying the
  setting to be configured.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkLayerSettingsCreateInfoEXT-sType-sType"
  id="VUID-VkLayerSettingsCreateInfoEXT-sType-sType"></a>
  VUID-VkLayerSettingsCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LAYER_SETTINGS_CREATE_INFO_EXT`

- <a href="#VUID-VkLayerSettingsCreateInfoEXT-pSettings-parameter"
  id="VUID-VkLayerSettingsCreateInfoEXT-pSettings-parameter"></a>
  VUID-VkLayerSettingsCreateInfoEXT-pSettings-parameter  
  If `settingCount` is not `0`, `pSettings` **must** be a valid pointer
  to an array of `settingCount` valid
  [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_layer_settings](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_layer_settings.html),
[VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLayerSettingsCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
