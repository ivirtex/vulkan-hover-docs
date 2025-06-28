# VkDeviceImageMemoryRequirements(3) Manual Page

## Name

VkDeviceImageMemoryRequirements - (None)



## [](#_c_specification)C Specification

The `VkDeviceImageMemoryRequirements` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkDeviceImageMemoryRequirements {
    VkStructureType             sType;
    const void*                 pNext;
    const VkImageCreateInfo*    pCreateInfo;
    VkImageAspectFlagBits       planeAspect;
} VkDeviceImageMemoryRequirements;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance4
typedef VkDeviceImageMemoryRequirements VkDeviceImageMemoryRequirementsKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pCreateInfo` is a pointer to a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure containing parameters affecting creation of the image to query.
- `planeAspect` is a [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) value specifying the aspect corresponding to the image plane to query. This parameter is ignored unless `pCreateInfo->tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, or `pCreateInfo->flags` has `VK_IMAGE_CREATE_DISJOINT_BIT` set.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06416)VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06416  
  The `pCreateInfo->pNext` chain **must** not contain a [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSwapchainCreateInfoKHR.html) structure
- [](#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06776)VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06776  
  The `pCreateInfo->pNext` chain **must** not contain a [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html) structure
- [](#VUID-VkDeviceImageMemoryRequirements-pNext-06996)VUID-VkDeviceImageMemoryRequirements-pNext-06996  
  Applications also **must** not call [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirements.html) with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) whose `pNext` chain includes a [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html) structure with non-zero `externalFormat`
- [](#VUID-VkDeviceImageMemoryRequirements-pNext-08962)VUID-VkDeviceImageMemoryRequirements-pNext-08962  
  Applications also **must** not call [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirements.html) with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) whose `pNext` chain includes a [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatQNX.html) structure with non-zero `externalFormat`
- [](#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06417)VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06417  
  If `pCreateInfo->format` specifies a *multi-planar* format and `pCreateInfo->flags` has `VK_IMAGE_CREATE_DISJOINT_BIT` set then `planeAspect` **must** not be `VK_IMAGE_ASPECT_NONE_KHR`
- [](#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06419)VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06419  
  If `pCreateInfo->flags` has `VK_IMAGE_CREATE_DISJOINT_BIT` set and if the `pCreateInfo->tiling` is `VK_IMAGE_TILING_LINEAR` or `VK_IMAGE_TILING_OPTIMAL`, then `planeAspect` **must** be a single valid [multi-planar aspect mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar-image-aspect) bit
- [](#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06420)VUID-VkDeviceImageMemoryRequirements-pCreateInfo-06420  
  If `pCreateInfo->tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then `planeAspect` **must** be a single valid *memory plane* for the image (that is, `aspectMask` **must** specify a plane index that is less than the [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount` associated with the image’s `format` and [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`)

Valid Usage (Implicit)

- [](#VUID-VkDeviceImageMemoryRequirements-sType-sType)VUID-VkDeviceImageMemoryRequirements-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS`
- [](#VUID-VkDeviceImageMemoryRequirements-pNext-pNext)VUID-VkDeviceImageMemoryRequirements-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDeviceImageMemoryRequirements-pCreateInfo-parameter)VUID-VkDeviceImageMemoryRequirements-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure
- [](#VUID-VkDeviceImageMemoryRequirements-planeAspect-parameter)VUID-VkDeviceImageMemoryRequirements-planeAspect-parameter  
  If `planeAspect` is not `0`, `planeAspect` **must** be a valid [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_maintenance4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance4.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirements.html), [vkGetDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageMemoryRequirementsKHR.html), [vkGetDeviceImageSparseMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSparseMemoryRequirements.html), [vkGetDeviceImageSparseMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSparseMemoryRequirementsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceImageMemoryRequirements)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0