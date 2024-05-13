# VkLayerSettingTypeEXT(3) Manual Page

## Name

VkLayerSettingTypeEXT - Type of the values that can be passed to a layer



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of [VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html)::`type`,
specifying the type of the data returned in
[VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html)::`pValues`, are:

``` c
// Provided by VK_EXT_layer_settings
typedef enum VkLayerSettingTypeEXT {
    VK_LAYER_SETTING_TYPE_BOOL32_EXT = 0,
    VK_LAYER_SETTING_TYPE_INT32_EXT = 1,
    VK_LAYER_SETTING_TYPE_INT64_EXT = 2,
    VK_LAYER_SETTING_TYPE_UINT32_EXT = 3,
    VK_LAYER_SETTING_TYPE_UINT64_EXT = 4,
    VK_LAYER_SETTING_TYPE_FLOAT32_EXT = 5,
    VK_LAYER_SETTING_TYPE_FLOAT64_EXT = 6,
    VK_LAYER_SETTING_TYPE_STRING_EXT = 7,
} VkLayerSettingTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_LAYER_SETTING_TYPE_BOOL32_EXT` specifies that the layer setting’s
  type is [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html).

- `VK_LAYER_SETTING_TYPE_INT32_EXT` specifies that the layer setting’s
  type is signed 32-bit integer.

- `VK_LAYER_SETTING_TYPE_INT64_EXT` specifies that the layer setting’s
  type is signed 64-bit integer.

- `VK_LAYER_SETTING_TYPE_UINT32_EXT` specifies that the layer setting’s
  type is unsigned 32-bit integer.

- `VK_LAYER_SETTING_TYPE_UINT64_EXT` specifies that the layer setting’s
  type is unsigned 64-bit integer.

- `VK_LAYER_SETTING_TYPE_FLOAT32_EXT` specifies that the layer setting’s
  type is 32-bit floating-point.

- `VK_LAYER_SETTING_TYPE_FLOAT64_EXT` specifies that the layer setting’s
  type is 64-bit floating-point.

- `VK_LAYER_SETTING_TYPE_STRING_EXT` specifies that the layer setting’s
  type is a pointer to a null-terminated UTF-8 string.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_layer_settings](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_layer_settings.html),
[VkLayerSettingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLayerSettingEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLayerSettingTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
