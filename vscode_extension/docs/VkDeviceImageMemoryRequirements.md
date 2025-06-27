# VkDeviceImageMemoryRequirements(3) Manual Page

## Name

VkDeviceImageMemoryRequirements - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceImageMemoryRequirements` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkDeviceImageMemoryRequirements {
    VkStructureType             sType;
    const void*                 pNext;
    const VkImageCreateInfo*    pCreateInfo;
    VkImageAspectFlagBits       planeAspect;
} VkDeviceImageMemoryRequirements;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance4
typedef VkDeviceImageMemoryRequirements VkDeviceImageMemoryRequirementsKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pCreateInfo` is a pointer to a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure containing
  parameters affecting creation of the image to query.

- `planeAspect` is a [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html)
  value specifying the aspect corresponding to the image plane to query.
  This parameter is ignored unless `pCreateInfo->tiling` is
  `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, or `pCreateInfo->flags` has
  `VK_IMAGE_CREATE_DISJOINT_BIT` set.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06416"
  id="VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06416"></a>
  VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06416  
  The `pCreateInfo->pNext` chain **must** not contain a
  [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html)
  structure

- <a href="#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06776"
  id="VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06776"></a>
  VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06776  
  The `pCreateInfo->pNext` chain **must** not contain a
  [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)
  structure

- <a href="#VUID-VkDeviceImageMemoryRequirements-pNext-06996"
  id="VUID-VkDeviceImageMemoryRequirements-pNext-06996"></a>
  VUID-VkDeviceImageMemoryRequirements-pNext-06996  
  Applications also **must** not call
  [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirements.html)
  with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) whose `pNext` chain
  includes a [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)
  structure with non-zero `externalFormat`

- <a href="#VUID-VkDeviceImageMemoryRequirements-pNext-08962"
  id="VUID-VkDeviceImageMemoryRequirements-pNext-08962"></a>
  VUID-VkDeviceImageMemoryRequirements-pNext-08962  
  Applications also **must** not call
  [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirements.html)
  with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) whose `pNext` chain
  includes a [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatQNX.html) structure
  with non-zero `externalFormat`

- <a href="#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06417"
  id="VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06417"></a>
  VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06417  
  If `pCreateInfo->format` specifies a *multi-planar* format and
  `pCreateInfo->flags` has `VK_IMAGE_CREATE_DISJOINT_BIT` set then
  `planeAspect` **must** not be `VK_IMAGE_ASPECT_NONE_KHR`

- <a href="#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06419"
  id="VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06419"></a>
  VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06419  
  If `pCreateInfo->flags` has `VK_IMAGE_CREATE_DISJOINT_BIT` set and if
  the `pCreateInfo->tiling` is `VK_IMAGE_TILING_LINEAR` or
  `VK_IMAGE_TILING_OPTIMAL`, then `planeAspect` **must** be a single
  valid <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-planes-image-aspect"
  target="_blank" rel="noopener">multi-planar aspect mask</a> bit

- <a href="#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06420"
  id="VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06420"></a>
  VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06420  
  If `pCreateInfo->tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`,
  then `planeAspect` **must** be a single valid *memory plane* for the
  image (that is, `aspectMask` **must** specify a plane index that is
  less than the
  [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount`
  associated with the image’s `format` and
  [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`)

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceImageMemoryRequirements-sType-sType"
  id="VUID-VkDeviceImageMemoryRequirements-sType-sType"></a>
  VUID-VkDeviceImageMemoryRequirements-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS`

- <a href="#VUID-VkDeviceImageMemoryRequirements-pNext-pNext"
  id="VUID-VkDeviceImageMemoryRequirements-pNext-pNext"></a>
  VUID-VkDeviceImageMemoryRequirements-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-parameter"
  id="VUID-VkDeviceImageMemoryRequirements-pCreateInfo-parameter"></a>
  VUID-VkDeviceImageMemoryRequirements-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure

- <a href="#VUID-VkDeviceImageMemoryRequirements-planeAspect-parameter"
  id="VUID-VkDeviceImageMemoryRequirements-planeAspect-parameter"></a>
  VUID-VkDeviceImageMemoryRequirements-planeAspect-parameter  
  If `planeAspect` is not `0`, `planeAspect` **must** be a valid
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance4.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirements.html),
[vkGetDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirementsKHR.html),
[vkGetDeviceImageSparseMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSparseMemoryRequirements.html),
[vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceImageMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
