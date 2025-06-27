# VkCommandBufferInheritanceRenderPassTransformInfoQCOM(3) Manual Page

## Name

VkCommandBufferInheritanceRenderPassTransformInfoQCOM - Structure
describing transformed render pass parameters command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin recording a secondary command buffer compatible with execution
inside a render pass using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-renderpass-transform"
target="_blank" rel="noopener">render pass transform</a>, add the
[VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html)
to the `pNext` chain of
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
structure passed to the
[vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html) command specifying the
parameters for transformed rasterization.

The `VkCommandBufferInheritanceRenderPassTransformInfoQCOM` structure is
defined as:

``` c
// Provided by VK_QCOM_render_pass_transform
typedef struct VkCommandBufferInheritanceRenderPassTransformInfoQCOM {
    VkStructureType                  sType;
    void*                            pNext;
    VkSurfaceTransformFlagBitsKHR    transform;
    VkRect2D                         renderArea;
} VkCommandBufferInheritanceRenderPassTransformInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `transform` is a
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value describing the transform to be applied to the render pass.

- `renderArea` is the render area that is affected by the command
  buffer.

## <a href="#_description" class="anchor"></a>Description

When the secondary is recorded to execute within a render pass instance
using [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html), the render pass
transform parameters of the secondary command buffer **must** be
consistent with the render pass transform parameters specified for the
render pass instance. In particular, the `transform` and `renderArea`
for command buffer **must** be identical to the `transform` and
`renderArea` of the render pass instance.

Valid Usage

- <a
  href="#VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-transform-02864"
  id="VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-transform-02864"></a>
  VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-transform-02864  
  `transform` **must** be `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`,
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`,
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`

Valid Usage (Implicit)

- <a
  href="#VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-sType-sType"
  id="VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-sType-sType"></a>
  VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDER_PASS_TRANSFORM_INFO_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_render_pass_transform](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_render_pass_transform.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferInheritanceRenderPassTransformInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
