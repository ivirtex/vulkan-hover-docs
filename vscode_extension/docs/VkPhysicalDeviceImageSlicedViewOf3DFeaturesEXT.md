# VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT - Structure describing
whether slice-based views of 3D images can be used in storage image
descriptors



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_image_sliced_view_of_3d
typedef struct VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           imageSlicedViewOf3D;
} VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT`
structure describe the following features:

## <a href="#_description" class="anchor"></a>Description

- <span id="features-imageSlicedViewOf3D"></span> `imageSlicedViewOf3D`
  indicates that the implementation supports using a sliced view of a 3D
  image in a descriptor of type `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` by
  using a
  [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSlicedCreateInfoEXT.html)
  structure when creating the view.

If the `VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_SLICED_VIEW_OF_3D_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_sliced_view_of_3d](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_sliced_view_of_3d.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
