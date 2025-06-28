# VkPhysicalDeviceImageDrmFormatModifierInfoEXT(3) Manual Page

## Name

VkPhysicalDeviceImageDrmFormatModifierInfoEXT - Structure specifying a DRM format modifier as image creation parameter



## [](#_c_specification)C Specification

To query the image capabilities that are compatible with a [Linux DRM format modifier](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-drm-format-modifier), set [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`tiling` to `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT` and add a [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html) structure to the `pNext` chain of [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html).

The [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkPhysicalDeviceImageDrmFormatModifierInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           drmFormatModifier;
    VkSharingMode      sharingMode;
    uint32_t           queueFamilyIndexCount;
    const uint32_t*    pQueueFamilyIndices;
} VkPhysicalDeviceImageDrmFormatModifierInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `drmFormatModifier` is the image’s *Linux DRM format modifier*, corresponding to [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)::`drmFormatModifier` or to [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)::`pDrmFormatModifiers`.
- `sharingMode` specifies how the image will be accessed by multiple queue families.
- `queueFamilyIndexCount` is the number of entries in the `pQueueFamilyIndices` array.
- `pQueueFamilyIndices` is a pointer to an array of queue families that will access the image. It is ignored if `sharingMode` is not `VK_SHARING_MODE_CONCURRENT`.

## [](#_description)Description

If the `drmFormatModifier` is incompatible with the parameters specified in [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) and its `pNext` chain, then [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) returns `VK_ERROR_FORMAT_NOT_SUPPORTED`. The implementation **must** support the query of any `drmFormatModifier`, including unknown and invalid modifier values.

Valid Usage

- [](#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02314)VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02314  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, then `pQueueFamilyIndices` **must** be a valid pointer to an array of `queueFamilyIndexCount` `uint32_t` values
- [](#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02315)VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02315  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, then `queueFamilyIndexCount` **must** be greater than `1`
- [](#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02316)VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02316  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of `pQueueFamilyIndices` **must** be unique and **must** be less than the `pQueueFamilyPropertyCount` returned by [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html) for the `physicalDevice` that was used to create `device`

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sType-sType)VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT`
- [](#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-parameter)VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-parameter  
  `sharingMode` **must** be a valid [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html) value

## [](#_see_also)See Also

[VK\_EXT\_image\_drm\_format\_modifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_drm_format_modifier.html), [VkSharingMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharingMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageDrmFormatModifierInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0