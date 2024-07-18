# VkPhysicalDeviceImageAlignmentControlPropertiesMESA(3) Manual Page

## Name

VkPhysicalDeviceImageAlignmentControlPropertiesMESA - Structure
describing supported image alignments for a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageAlignmentControlPropertiesMESA` structure is
defined as:

``` c
// Provided by VK_MESA_image_alignment_control
typedef struct VkPhysicalDeviceImageAlignmentControlPropertiesMESA {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           supportedImageAlignmentMask;
} VkPhysicalDeviceImageAlignmentControlPropertiesMESA;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceImageAlignmentControlPropertiesMESA`
structure describe the following:

## <a href="#_description" class="anchor"></a>Description

- <span id="limits-supportedImageAlignmentMask"></span>
  `supportedImageAlignmentMask` is a bitwise-or of all potentially
  supported image alignments for a given physical device when using
  `VK_IMAGE_TILING_OPTIMAL`. If a given alignment is supported, the
  application **can** request an image to have that alignment. A given
  set of image creation parameters **may** support a subset of these
  alignments. To determine if a particular alignment is supported for a
  given set of image creation parameters, check
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)::`alignment` after
  chaining in
  [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAlignmentControlCreateInfoMESA.html).

If the `VkPhysicalDeviceImageAlignmentControlPropertiesMESA` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceImageAlignmentControlPropertiesMESA-sType-sType"
  id="VUID-VkPhysicalDeviceImageAlignmentControlPropertiesMESA-sType-sType"></a>
  VUID-VkPhysicalDeviceImageAlignmentControlPropertiesMESA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ALIGNMENT_CONTROL_PROPERTIES_MESA`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_MESA_image_alignment_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MESA_image_alignment_control.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageAlignmentControlPropertiesMESA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
