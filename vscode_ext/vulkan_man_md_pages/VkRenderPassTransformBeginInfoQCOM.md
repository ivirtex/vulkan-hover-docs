# VkRenderPassTransformBeginInfoQCOM(3) Manual Page

## Name

VkRenderPassTransformBeginInfoQCOM - Structure describing transform
parameters of a render pass instance



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin a render pass instance with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-renderpass-transform"
target="_blank" rel="noopener">render pass transform</a> enabled, add
the
[VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)
to the `pNext` chain of
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) structure passed to
the [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html) command specifying
the render pass transform.

The `VkRenderPassTransformBeginInfoQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_render_pass_transform
typedef struct VkRenderPassTransformBeginInfoQCOM {
    VkStructureType                  sType;
    void*                            pNext;
    VkSurfaceTransformFlagBitsKHR    transform;
} VkRenderPassTransformBeginInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `transform` is a
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value describing the transform to be applied to rasterization.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkRenderPassTransformBeginInfoQCOM-transform-02871"
  id="VUID-VkRenderPassTransformBeginInfoQCOM-transform-02871"></a>
  VUID-VkRenderPassTransformBeginInfoQCOM-transform-02871  
  `transform` **must** be `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`,
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`,
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`

- <a href="#VUID-VkRenderPassTransformBeginInfoQCOM-flags-02872"
  id="VUID-VkRenderPassTransformBeginInfoQCOM-flags-02872"></a>
  VUID-VkRenderPassTransformBeginInfoQCOM-flags-02872  
  The `renderpass` **must** have been created with
  [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html)::`flags`
  containing `VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassTransformBeginInfoQCOM-sType-sType"
  id="VUID-VkRenderPassTransformBeginInfoQCOM-sType-sType"></a>
  VUID-VkRenderPassTransformBeginInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_TRANSFORM_BEGIN_INFO_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_render_pass_transform](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_render_pass_transform.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassTransformBeginInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
