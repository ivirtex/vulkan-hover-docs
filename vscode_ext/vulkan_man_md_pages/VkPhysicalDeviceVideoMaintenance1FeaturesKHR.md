# VkPhysicalDeviceVideoMaintenance1FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoMaintenance1FeaturesKHR - Structure describing the
video maintenance features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVideoMaintenance1FeaturesKHR` structure is defined
as:

``` c
// Provided by VK_KHR_video_maintenance1
typedef struct VkPhysicalDeviceVideoMaintenance1FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           videoMaintenance1;
} VkPhysicalDeviceVideoMaintenance1FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-videoMaintenance1"></span> `videoMaintenance1`
  indicates that the implementation supports the following:

  - The new buffer creation flag
    `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`.

  - The new image creation flag
    `VK_IMAGE_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR`.

  - The new video session creation flag
    `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceVideoMaintenance1FeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceVideoMaintenance1FeaturesKHR` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceVideoMaintenance1FeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceVideoMaintenance1FeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceVideoMaintenance1FeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_MAINTENANCE_1_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_maintenance1.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVideoMaintenance1FeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
