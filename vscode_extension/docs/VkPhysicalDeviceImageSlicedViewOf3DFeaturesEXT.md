# VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT - Structure describing whether slice-based views of 3D images can be used in storage image descriptors



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_image_sliced_view_of_3d
typedef struct VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           imageSlicedViewOf3D;
} VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT` structure describe the following features:

## [](#_description)Description

- []()`imageSlicedViewOf3D` indicates that the implementation supports using a sliced view of a 3D image in a descriptor of type `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` by using a [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSlicedCreateInfoEXT.html) structure when creating the view.

If the `VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_SLICED_VIEW_OF_3D_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_image\_sliced\_view\_of\_3d](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_sliced_view_of_3d.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0