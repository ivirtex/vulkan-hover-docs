# VkCopyCommandTransformInfoQCOM(3) Manual Page

## Name

VkCopyCommandTransformInfoQCOM - Structure describing transform parameters of rotated copy command



## [](#_c_specification)C Specification

The `VkCopyCommandTransformInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_rotated_copy_commands
typedef struct VkCopyCommandTransformInfoQCOM {
    VkStructureType                  sType;
    const void*                      pNext;
    VkSurfaceTransformFlagBitsKHR    transform;
} VkCopyCommandTransformInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `transform` is a [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html) value describing the transform to be applied.

## [](#_description)Description

Including this structure in the `pNext` chain of [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html) defines a rotation to be performed when copying between an image and a buffer. Including this structure in the `pNext` chain of [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageInfo2.html) defines a rotation to be performed when blitting between two images. If this structure is not specified in either case, the implementation behaves as if it was specified with a `transform` equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`.

Specifying a transform for a copy between an image and a buffer [rotates the region accessed in the image around the offset](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers-images-rotation-addressing). Specifying a transform for a blit performs a similar transform as described in [Image Blits with Scaling and Rotation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-images-scaling-rotation).

Rotations other than `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` **can** only be specified for single-plane 2D images with a 1x1x1 [texel block extent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes).

Valid Usage

- [](#VUID-VkCopyCommandTransformInfoQCOM-transform-04560)VUID-VkCopyCommandTransformInfoQCOM-transform-04560  
  `transform` **must** be `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`, `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, or `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-VkCopyCommandTransformInfoQCOM-sType-sType)VUID-VkCopyCommandTransformInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_COMMAND_TRANSFORM_INFO_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_rotated\_copy\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_rotated_copy_commands.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceTransformFlagBitsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyCommandTransformInfoQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0