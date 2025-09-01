# vkGetRenderAreaGranularity(3) Manual Page

## Name

vkGetRenderAreaGranularity - Returns the granularity for optimal render area



## [](#_c_specification)C Specification

To query the render area granularity, call:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkGetRenderAreaGranularity(
    VkDevice                                    device,
    VkRenderPass                                renderPass,
    VkExtent2D*                                 pGranularity);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the render pass.
- `renderPass` is a handle to a render pass.
- `pGranularity` is a pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) structure in which the granularity is returned.

## [](#_description)Description

The conditions leading to an optimal `renderArea` are:

- the `offset.x` member in `renderArea` is a multiple of the `width` member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) (the horizontal granularity).
- the `offset.y` member in `renderArea` is a multiple of the `height` member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) (the vertical granularity).
- either the `extent.width` member in `renderArea` is a multiple of the horizontal granularity or `offset.x`+`extent.width` is equal to the `width` of the `framebuffer` in the [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html).
- either the `extent.height` member in `renderArea` is a multiple of the vertical granularity or `offset.y`+`extent.height` is equal to the `height` of the `framebuffer` in the [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html).

Subpass dependencies are not affected by the render area, and apply to the entire image subresources attached to the framebuffer as specified in the description of [automatic layout transitions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-layout-transitions). Similarly, pipeline barriers are valid even if their effect extends outside the render area.

Valid Usage (Implicit)

- [](#VUID-vkGetRenderAreaGranularity-device-parameter)VUID-vkGetRenderAreaGranularity-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetRenderAreaGranularity-renderPass-parameter)VUID-vkGetRenderAreaGranularity-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) handle
- [](#VUID-vkGetRenderAreaGranularity-pGranularity-parameter)VUID-vkGetRenderAreaGranularity-pGranularity-parameter  
  `pGranularity` **must** be a valid pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) structure
- [](#VUID-vkGetRenderAreaGranularity-renderPass-parent)VUID-vkGetRenderAreaGranularity-renderPass-parent  
  `renderPass` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetRenderAreaGranularity).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0