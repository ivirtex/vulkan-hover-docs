# VkPhysicalDeviceImageDrmFormatModifierInfoEXT(3) Manual Page

## Name

VkPhysicalDeviceImageDrmFormatModifierInfoEXT - Structure specifying a
DRM format modifier as image creation parameter



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the image capabilities that are compatible with a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">Linux DRM format modifier</a>, set
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`tiling`
to `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT` and add a
[VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)
structure to the `pNext` chain of
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html).

The
[VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)
structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `drmFormatModifier` is the image’s *Linux DRM format modifier*,
  corresponding to
  [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)::`modifier`
  or to
  [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)::`pModifiers`.

- `sharingMode` specifies how the image will be accessed by multiple
  queue families.

- `queueFamilyIndexCount` is the number of entries in the
  `pQueueFamilyIndices` array.

- `pQueueFamilyIndices` is a pointer to an array of queue families that
  will access the image. It is ignored if `sharingMode` is not
  `VK_SHARING_MODE_CONCURRENT`.

## <a href="#_description" class="anchor"></a>Description

If the `drmFormatModifier` is incompatible with the parameters specified
in
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
and its `pNext` chain, then
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
returns `VK_ERROR_FORMAT_NOT_SUPPORTED`. The implementation **must**
support the query of any `drmFormatModifier`, including unknown and
invalid modifier values.

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02314"
  id="VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02314"></a>
  VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02314  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, then
  `pQueueFamilyIndices` **must** be a valid pointer to an array of
  `queueFamilyIndexCount` `uint32_t` values

- <a
  href="#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02315"
  id="VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02315"></a>
  VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02315  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, then
  `queueFamilyIndexCount` **must** be greater than `1`

- <a
  href="#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02316"
  id="VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02316"></a>
  VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-02316  
  If `sharingMode` is `VK_SHARING_MODE_CONCURRENT`, each element of
  `pQueueFamilyIndices` **must** be unique and **must** be less than the
  `pQueueFamilyPropertyCount` returned by
  [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html)
  for the `physicalDevice` that was used to create `device`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sType-sType"
  id="VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT`

- <a
  href="#VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-parameter"
  id="VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-parameter"></a>
  VUID-VkPhysicalDeviceImageDrmFormatModifierInfoEXT-sharingMode-parameter  
  `sharingMode` **must** be a valid [VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VkSharingMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharingMode.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageDrmFormatModifierInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
