# VkImageAlignmentControlCreateInfoMESA(3) Manual Page

## Name

VkImageAlignmentControlCreateInfoMESA - Specify image alignment



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` list of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
includes a `VkImageAlignmentControlCreateInfoMESA` structure, then that
structure describes desired alignment for this image.

The `VkImageAlignmentControlCreateInfoMESA` structure is defined as:

``` c
// Provided by VK_MESA_image_alignment_control
typedef struct VkImageAlignmentControlCreateInfoMESA {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           maximumRequestedAlignment;
} VkImageAlignmentControlCreateInfoMESA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maximumRequestedAlignment` specifies the maximum alignment for the
  image.

## <a href="#_description" class="anchor"></a>Description

If `maximumRequestedAlignment` is not 0, the implementation **should**
choose an image memory layout that requires an alignment no larger than
`maximumRequestedAlignment` as reported in
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)::`alignment`. If such
a layout does not exist for the given image creation parameters, the
implementation **should** return the smallest alignment which is
supported in [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html).

If an implementation needs to disable image compression for
`maximumRequestedAlignment` to be honored - where a larger alignment
would enable image compression - the implementation **should** not use
`maximumRequestedAlignment`, and **should** return the smallest
alignment which does not compromise compression. If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imageCompressionControl"
target="_blank" rel="noopener"><code>imageCompressionControl</code></a>
is enabled, the application **can** chain a
[VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionControlEXT.html) with
`VK_IMAGE_COMPRESSION_DISABLED_EXT`. In this case, image compression
considerations **should** not apply when implementation decides
alignment.

Valid Usage

- <a
  href="#VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09655"
  id="VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09655"></a>
  VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09655  
  If `maximumRequestedAlignment` is not 0, `maximumRequestedAlignment`
  **must** be a power of two

- <a
  href="#VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09656"
  id="VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09656"></a>
  VUID-VkImageAlignmentControlCreateInfoMESA-maximumRequestedAlignment-09656  
  If `maximumRequestedAlignment` is not 0, the bitwise-and of
  `maximumRequestedAlignment` and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-supportedImageAlignmentMask"
  target="_blank"
  rel="noopener"><code>supportedImageAlignmentMask</code></a> **must**
  be non-zero

- <a
  href="#VUID-VkImageAlignmentControlCreateInfoMESA-imageAlignmentControl-09657"
  id="VUID-VkImageAlignmentControlCreateInfoMESA-imageAlignmentControl-09657"></a>
  VUID-VkImageAlignmentControlCreateInfoMESA-imageAlignmentControl-09657  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imageAlignmentControl"
  target="_blank" rel="noopener"><code>imageAlignmentControl</code></a>
  **must** be enabled on the device

Valid Usage (Implicit)

- <a href="#VUID-VkImageAlignmentControlCreateInfoMESA-sType-sType"
  id="VUID-VkImageAlignmentControlCreateInfoMESA-sType-sType"></a>
  VUID-VkImageAlignmentControlCreateInfoMESA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_ALIGNMENT_CONTROL_CREATE_INFO_MESA`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_MESA_image_alignment_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MESA_image_alignment_control.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageAlignmentControlCreateInfoMESA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
