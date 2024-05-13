# VkLayerSettingEXT(3) Manual Page

## Name

VkLayerSettingEXT - Specify a layer capability to configure



## <a href="#_c_specification" class="anchor"></a>C Specification

The values of elements of the
[VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html)::`pSettings`
array, specifying layer settings to be configured, are:

``` c
// Provided by VK_EXT_layer_settings
typedef struct VkLayerSettingEXT {
    const char*              pLayerName;
    const char*              pSettingName;
    VkLayerSettingTypeEXT    type;
    uint32_t                 valueCount;
    const void*              pValues;
} VkLayerSettingEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `pLayerName` is a pointer to a null-terminated UTF-8 string naming the
  layer to configure the setting from.

- `pSettingName` is a pointer to a null-terminated UTF-8 string naming
  the setting to configure. Unknown `pSettingName` by the layer are
  ignored.

- `type` is a [VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingTypeEXT.html) value
  specifying the type of the `pValues` values.

- `count` is the number of values used to configure the layer setting.

- `pValues` is a pointer to an array of `count` values of the type
  indicated by `type` to configure the layer setting.

## <a href="#_description" class="anchor"></a>Description

When multiple
[VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html)
structures are chained and the same `pSettingName` is referenced for the
same `pLayerName`, the value of the first reference of the layer setting
is used.

Valid Usage (Implicit)

- <a href="#VUID-VkLayerSettingEXT-pLayerName-parameter"
  id="VUID-VkLayerSettingEXT-pLayerName-parameter"></a>
  VUID-VkLayerSettingEXT-pLayerName-parameter  
  `pLayerName` **must** be a null-terminated UTF-8 string

- <a href="#VUID-VkLayerSettingEXT-pSettingName-parameter"
  id="VUID-VkLayerSettingEXT-pSettingName-parameter"></a>
  VUID-VkLayerSettingEXT-pSettingName-parameter  
  `pSettingName` **must** be a null-terminated UTF-8 string

- <a href="#VUID-VkLayerSettingEXT-type-parameter"
  id="VUID-VkLayerSettingEXT-type-parameter"></a>
  VUID-VkLayerSettingEXT-type-parameter  
  `type` **must** be a valid
  [VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingTypeEXT.html) value

- <a href="#VUID-VkLayerSettingEXT-pValues-parameter"
  id="VUID-VkLayerSettingEXT-pValues-parameter"></a>
  VUID-VkLayerSettingEXT-pValues-parameter  
  If `valueCount` is not `0`, `pValues` **must** be a valid pointer to
  an array of `valueCount` bytes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_layer_settings](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_layer_settings.html),
[VkLayerSettingTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingTypeEXT.html),
[VkLayerSettingsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingsCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLayerSettingEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
