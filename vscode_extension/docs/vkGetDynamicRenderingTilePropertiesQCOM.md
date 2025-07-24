# vkGetDynamicRenderingTilePropertiesQCOM(3) Manual Page

## Name

vkGetDynamicRenderingTilePropertiesQCOM - Get the properties when using dynamic rendering



## [](#_c_specification)C Specification

To query the tile properties when using dynamic rendering, call:

```c++
// Provided by VK_QCOM_tile_properties
VkResult vkGetDynamicRenderingTilePropertiesQCOM(
    VkDevice                                    device,
    const VkRenderingInfo*                      pRenderingInfo,
    VkTilePropertiesQCOM*                       pProperties);
```

## [](#_parameters)Parameters

- `device` is a logical device associated with the render pass.
- `pRenderingInfo` is a pointer to the [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) structure specifying details of the render pass instance in dynamic rendering.
- `pProperties` is a pointer to a [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTilePropertiesQCOM.html) structure in which the properties are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDynamicRenderingTilePropertiesQCOM-device-parameter)VUID-vkGetDynamicRenderingTilePropertiesQCOM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDynamicRenderingTilePropertiesQCOM-pRenderingInfo-parameter)VUID-vkGetDynamicRenderingTilePropertiesQCOM-pRenderingInfo-parameter  
  `pRenderingInfo` **must** be a valid pointer to a valid [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) structure
- [](#VUID-vkGetDynamicRenderingTilePropertiesQCOM-pProperties-parameter)VUID-vkGetDynamicRenderingTilePropertiesQCOM-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTilePropertiesQCOM.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html), [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTilePropertiesQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDynamicRenderingTilePropertiesQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0