# VkCommandBufferInheritanceRenderPassTransformInfoQCOM(3) Manual Page

## Name

VkCommandBufferInheritanceRenderPassTransformInfoQCOM - Structure describing transformed render pass parameters command buffer



## [](#_c_specification)C Specification

To begin recording a secondary command buffer compatible with execution inside a render pass using [render pass transform](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-renderpass-transform), add the [VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html) to the `pNext` chain of [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) structure passed to the [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBeginCommandBuffer.html) command specifying the parameters for transformed rasterization.

The `VkCommandBufferInheritanceRenderPassTransformInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_render_pass_transform
typedef struct VkCommandBufferInheritanceRenderPassTransformInfoQCOM {
    VkStructureType                  sType;
    void*                            pNext;
    VkSurfaceTransformFlagBitsKHR    transform;
    VkRect2D                         renderArea;
} VkCommandBufferInheritanceRenderPassTransformInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `transform` is a [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html) value describing the transform to be applied to the render pass.
- `renderArea` is the render area that is affected by the command buffer.

## [](#_description)Description

When the secondary is recorded to execute within a render pass instance using [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteCommands.html), the render pass transform parameters of the secondary command buffer **must** be consistent with the render pass transform parameters specified for the render pass instance. In particular, the `transform` and `renderArea` for command buffer **must** be identical to the `transform` and `renderArea` of the render pass instance.

Valid Usage

- [](#VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-transform-02864)VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-transform-02864  
  `transform` **must** be `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-sType-sType)VUID-VkCommandBufferInheritanceRenderPassTransformInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDER_PASS_TRANSFORM_INFO_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_render\_pass\_transform](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_render_pass_transform.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferInheritanceRenderPassTransformInfoQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0