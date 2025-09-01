# vkGetRenderingAreaGranularity(3) Manual Page

## Name

vkGetRenderingAreaGranularity - Returns the granularity for dynamic rendering optimal render area



## [](#_c_specification)C Specification

To query the render area granularity for a render pass instance, call:

```c++
// Provided by VK_VERSION_1_4
void vkGetRenderingAreaGranularity(
    VkDevice                                    device,
    const VkRenderingAreaInfo*                  pRenderingAreaInfo,
    VkExtent2D*                                 pGranularity);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance5
void vkGetRenderingAreaGranularityKHR(
    VkDevice                                    device,
    const VkRenderingAreaInfo*                  pRenderingAreaInfo,
    VkExtent2D*                                 pGranularity);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the render pass instance.
- `pRenderingAreaInfo` is a pointer to a [VkRenderingAreaInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAreaInfo.html) structure specifying details of the render pass instance to query the render area granularity for.
- `pGranularity` is a pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) structure in which the granularity is returned.

## [](#_description)Description

The conditions leading to an optimal `renderArea` are:

- the `offset.x` member in `renderArea` is a multiple of the `width` member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) (the horizontal granularity).
- the `offset.y` member in `renderArea` is a multiple of the `height` member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) (the vertical granularity).
- either the `extent.width` member in `renderArea` is a multiple of the horizontal granularity or `offset.x`+`extent.width` is equal to the `width` of each attachment used in the render pass instance.
- either the `extent.height` member in `renderArea` is a multiple of the vertical granularity or `offset.y`+`extent.height` is equal to the `height` of each attachment used in the render pass instance.

Valid Usage (Implicit)

- [](#VUID-vkGetRenderingAreaGranularity-device-parameter)VUID-vkGetRenderingAreaGranularity-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetRenderingAreaGranularity-pRenderingAreaInfo-parameter)VUID-vkGetRenderingAreaGranularity-pRenderingAreaInfo-parameter  
  `pRenderingAreaInfo` **must** be a valid pointer to a valid [VkRenderingAreaInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAreaInfo.html) structure
- [](#VUID-vkGetRenderingAreaGranularity-pGranularity-parameter)VUID-vkGetRenderingAreaGranularity-pGranularity-parameter  
  `pGranularity` **must** be a valid pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) structure

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkRenderingAreaInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAreaInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetRenderingAreaGranularity).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0