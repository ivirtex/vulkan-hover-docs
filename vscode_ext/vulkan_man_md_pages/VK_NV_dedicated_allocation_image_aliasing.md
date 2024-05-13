# VK_NV_dedicated_allocation_image_aliasing(3) Manual Page

## Name

VK_NV_dedicated_allocation_image_aliasing - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

241

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

     [VK_KHR_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html)  
     and  
    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Nuno Subtil <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_dedicated_allocation_image_aliasing%5D%20@nsubtil%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_dedicated_allocation_image_aliasing%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>nsubtil</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-01-04

**Contributors**  
- Nuno Subtil, NVIDIA

- Jeff Bolz, NVIDIA

- Eric Werness, NVIDIA

- Axel Gneiting, id Software

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to alias images on dedicated
allocations, subject to specific restrictions: the extent and the number
of layers in the image being aliased must be smaller than or equal to
those of the original image for which the allocation was created, and
every other image parameter must match.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_EXTENSION_NAME`

- `VK_NV_DEDICATED_ALLOCATION_IMAGE_ALIASING_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEDICATED_ALLOCATION_IMAGE_ALIASING_FEATURES_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-01-04 (Nuno Subtil)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_dedicated_allocation_image_aliasing"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
