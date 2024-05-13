# VkRect2D(3) Manual Page

## Name

VkRect2D - Structure specifying a two-dimensional subregion



## <a href="#_c_specification" class="anchor"></a>C Specification

Rectangles are used to describe a specified rectangular region of pixels
within an image or framebuffer. Rectangles include both an offset and an
extent of the same dimensionality, as described above. Two-dimensional
rectangles are defined by the structure

``` c
// Provided by VK_VERSION_1_0
typedef struct VkRect2D {
    VkOffset2D    offset;
    VkExtent2D    extent;
} VkRect2D;
```

## <a href="#_members" class="anchor"></a>Members

- `offset` is a [VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html) specifying the rectangle
  offset.

- `extent` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) specifying the rectangle
  extent.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html),
[VkClearRect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearRect.html),
[VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html),
[VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html),
[VkDisplayPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPresentInfoKHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html),
[VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html),
[VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteInfoNV.html),
[VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html),
[VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html),
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html),
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
[VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeInfoARM.html),
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html),
[vkCmdSetDiscardRectangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEXT.html),
[vkCmdSetExclusiveScissorNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorNV.html),
[vkCmdSetScissor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissor.html),
[vkCmdSetScissorWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCount.html),
[vkCmdSetScissorWithCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCountEXT.html),
[vkGetPhysicalDevicePresentRectanglesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRect2D"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
