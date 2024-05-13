# vkGetRenderAreaGranularity(3) Manual Page

## Name

vkGetRenderAreaGranularity - Returns the granularity for optimal render
area



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the render area granularity, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetRenderAreaGranularity(
    VkDevice                                    device,
    VkRenderPass                                renderPass,
    VkExtent2D*                                 pGranularity);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the render pass.

- `renderPass` is a handle to a render pass.

- `pGranularity` is a pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html)
  structure in which the granularity is returned.

## <a href="#_description" class="anchor"></a>Description

The conditions leading to an optimal `renderArea` are:

- the `offset.x` member in `renderArea` is a multiple of the `width`
  member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) (the horizontal
  granularity).

- the `offset.y` member in `renderArea` is a multiple of the `height`
  member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) (the vertical
  granularity).

- either the `extent.width` member in `renderArea` is a multiple of the
  horizontal granularity or `offset.x`+`extent.width` is equal to the
  `width` of the `framebuffer` in the
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html).

- either the `extent.height` member in `renderArea` is a multiple of the
  vertical granularity or `offset.y`+`extent.height` is equal to the
  `height` of the `framebuffer` in the
  [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html).

Subpass dependencies are not affected by the render area, and apply to
the entire image subresources attached to the framebuffer as specified
in the description of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-layout-transitions"
target="_blank" rel="noopener">automatic layout transitions</a>.
Similarly, pipeline barriers are valid even if their effect extends
outside the render area.

Valid Usage (Implicit)

- <a href="#VUID-vkGetRenderAreaGranularity-device-parameter"
  id="VUID-vkGetRenderAreaGranularity-device-parameter"></a>
  VUID-vkGetRenderAreaGranularity-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetRenderAreaGranularity-renderPass-parameter"
  id="VUID-vkGetRenderAreaGranularity-renderPass-parameter"></a>
  VUID-vkGetRenderAreaGranularity-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

- <a href="#VUID-vkGetRenderAreaGranularity-pGranularity-parameter"
  id="VUID-vkGetRenderAreaGranularity-pGranularity-parameter"></a>
  VUID-vkGetRenderAreaGranularity-pGranularity-parameter  
  `pGranularity` **must** be a valid pointer to a
  [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) structure

- <a href="#VUID-vkGetRenderAreaGranularity-renderPass-parent"
  id="VUID-vkGetRenderAreaGranularity-renderPass-parent"></a>
  VUID-vkGetRenderAreaGranularity-renderPass-parent  
  `renderPass` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetRenderAreaGranularity"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
