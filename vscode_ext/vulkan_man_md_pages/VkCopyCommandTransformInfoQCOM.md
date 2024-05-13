# VkCopyCommandTransformInfoQCOM(3) Manual Page

## Name

VkCopyCommandTransformInfoQCOM - Structure describing transform
parameters of rotated copy command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyCommandTransformInfoQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_rotated_copy_commands
typedef struct VkCopyCommandTransformInfoQCOM {
    VkStructureType                  sType;
    const void*                      pNext;
    VkSurfaceTransformFlagBitsKHR    transform;
} VkCopyCommandTransformInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `transform` is a
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value describing the transform to be applied.

## <a href="#_description" class="anchor"></a>Description

Including this structure in the `pNext` chain of
[VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html) defines a rotation to be
performed when copying between an image and a buffer. Including this
structure in the `pNext` chain of
[VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html) defines a rotation to be
performed when blitting between two images. If this structure is not
specified in either case, the implementation behaves as if it was
specified with a `transform` equal to
`VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`.

Specifying a transform for a copy between an image and a buffer <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers-images-rotation-addressing"
target="_blank" rel="noopener">rotates the region accessed in the image
around the offset</a>. Specifying a transform for a blit performs a
similar transform as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-images-scaling-rotation"
target="_blank" rel="noopener">Image Blits with Scaling and Rotation</a>.

Rotations other than `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` **can**
only be specified for single-plane 2D images with a 1x1x1 <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility-classes"
target="_blank" rel="noopener">texel block extent</a>.

Valid Usage

- <a href="#VUID-VkCopyCommandTransformInfoQCOM-transform-04560"
  id="VUID-VkCopyCommandTransformInfoQCOM-transform-04560"></a>
  VUID-VkCopyCommandTransformInfoQCOM-transform-04560  
  `transform` **must** be `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`,
  `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR`,
  `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR`, or
  `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyCommandTransformInfoQCOM-sType-sType"
  id="VUID-VkCopyCommandTransformInfoQCOM-sType-sType"></a>
  VUID-VkCopyCommandTransformInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COPY_COMMAND_TRANSFORM_INFO_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_rotated_copy_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_rotated_copy_commands.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyCommandTransformInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
