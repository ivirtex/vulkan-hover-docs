# vkGetDynamicRenderingTilePropertiesQCOM(3) Manual Page

## Name

vkGetDynamicRenderingTilePropertiesQCOM - Get the properties when using
dynamic rendering



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the tile properties when using dynamic rendering, call:

``` c
// Provided by VK_QCOM_tile_properties
VkResult vkGetDynamicRenderingTilePropertiesQCOM(
    VkDevice                                    device,
    const VkRenderingInfo*                      pRenderingInfo,
    VkTilePropertiesQCOM*                       pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is a logical device associated with the render pass.

- `pRenderingInfo` is a pointer to the
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) structure specifying details
  of the render pass instance in dynamic rendering.

- `pProperties` is a pointer to a
  [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html) structure in which
  the properties are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetDynamicRenderingTilePropertiesQCOM-device-parameter"
  id="VUID-vkGetDynamicRenderingTilePropertiesQCOM-device-parameter"></a>
  VUID-vkGetDynamicRenderingTilePropertiesQCOM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDynamicRenderingTilePropertiesQCOM-pRenderingInfo-parameter"
  id="VUID-vkGetDynamicRenderingTilePropertiesQCOM-pRenderingInfo-parameter"></a>
  VUID-vkGetDynamicRenderingTilePropertiesQCOM-pRenderingInfo-parameter  
  `pRenderingInfo` **must** be a valid pointer to a valid
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) structure

- <a
  href="#VUID-vkGetDynamicRenderingTilePropertiesQCOM-pProperties-parameter"
  id="VUID-vkGetDynamicRenderingTilePropertiesQCOM-pProperties-parameter"></a>
  VUID-vkGetDynamicRenderingTilePropertiesQCOM-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a
  [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_tile_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_tile_properties.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html),
[VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDynamicRenderingTilePropertiesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
