# VkRenderPassTransformBeginInfoQCOM(3) Manual Page

## Name

VkRenderPassTransformBeginInfoQCOM - Structure describing transform parameters of a render pass instance



## [](#_c_specification)C Specification

To begin a render pass instance with [render pass transform](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-renderpass-transform) enabled, add the [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTransformBeginInfoQCOM.html) to the `pNext` chain of [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) structure passed to the [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass.html) command specifying the render pass transform.

The `VkRenderPassTransformBeginInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_render_pass_transform
typedef struct VkRenderPassTransformBeginInfoQCOM {
    VkStructureType                  sType;
    const void*                      pNext;
    VkSurfaceTransformFlagBitsKHR    transform;
} VkRenderPassTransformBeginInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `transform` is a [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html) value describing the transform to be applied to rasterization.

## [](#_description)Description

Valid Usage

- [](#VUID-VkRenderPassTransformBeginInfoQCOM-transform-02871)VUID-VkRenderPassTransformBeginInfoQCOM-transform-02871  
  `transform` **must** be `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`
- [](#VUID-VkRenderPassTransformBeginInfoQCOM-flags-02872)VUID-VkRenderPassTransformBeginInfoQCOM-flags-02872  
  The `renderpass` **must** have been created with [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html)::`flags` containing `VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM`

Valid Usage (Implicit)

- [](#VUID-VkRenderPassTransformBeginInfoQCOM-sType-sType)VUID-VkRenderPassTransformBeginInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_TRANSFORM_BEGIN_INFO_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_render\_pass\_transform](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_render_pass_transform.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassTransformBeginInfoQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0