# VkLayerSettingEXT(3) Manual Page

## Name

VkLayerSettingEXT - Specify a layer capability to configure



## [](#_c_specification)C Specification

The values of elements of the [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html)::`pSettings` array, specifying layer settings to be configured, are:

```c++
// Provided by VK_EXT_layer_settings
typedef struct VkLayerSettingEXT {
    const char*              pLayerName;
    const char*              pSettingName;
    VkLayerSettingTypeEXT    type;
    uint32_t                 valueCount;
    const void*              pValues;
} VkLayerSettingEXT;
```

## [](#_members)Members

- `pLayerName` is a pointer to a null-terminated UTF-8 string naming the layer to configure the setting from.
- `pSettingName` is a pointer to a null-terminated UTF-8 string naming the setting to configure. Values of `pSettingName` that are unknown to the layer are ignored.
- `type` is a [VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingTypeEXT.html) value specifying the type of the `pValues` values.
- `valueCount` is the number of values used to configure the layer setting.
- `pValues` is a pointer to an array of `valueCount` values of the type indicated by `type` to configure the layer setting.

## [](#_description)Description

When multiple [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html) structures are chained and the same `pSettingName` is referenced for the same `pLayerName`, the value of the first reference of the layer setting is used.

Valid Usage

- [](#VUID-VkLayerSettingEXT-valueCount-10070)VUID-VkLayerSettingEXT-valueCount-10070  
  If `valueCount` is not `0`, `pValues` **must** be a valid pointer to an array of `valueCount` values of the type indicated by `type`

Valid Usage (Implicit)

- [](#VUID-VkLayerSettingEXT-pLayerName-parameter)VUID-VkLayerSettingEXT-pLayerName-parameter  
  `pLayerName` **must** be a null-terminated UTF-8 string
- [](#VUID-VkLayerSettingEXT-pSettingName-parameter)VUID-VkLayerSettingEXT-pSettingName-parameter  
  `pSettingName` **must** be a null-terminated UTF-8 string
- [](#VUID-VkLayerSettingEXT-type-parameter)VUID-VkLayerSettingEXT-type-parameter  
  `type` **must** be a valid [VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingTypeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_layer\_settings](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_layer_settings.html), [VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingTypeEXT.html), [VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerSettingsCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLayerSettingEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0