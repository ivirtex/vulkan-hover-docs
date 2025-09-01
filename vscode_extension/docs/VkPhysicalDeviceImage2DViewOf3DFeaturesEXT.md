# VkPhysicalDeviceImage2DViewOf3DFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceImage2DViewOf3DFeaturesEXT - Structure describing whether single-slice 2D views of 3D images can be used in image descriptors



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImage2DViewOf3DFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_image_2d_view_of_3d
typedef struct VkPhysicalDeviceImage2DViewOf3DFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           image2DViewOf3D;
    VkBool32           sampler2DViewOf3D;
} VkPhysicalDeviceImage2DViewOf3DFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`image2DViewOf3D` indicates that the implementation supports using a 2D view of a 3D image in a descriptor of type `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` if the image is created using `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT`.
- []()`sampler2DViewOf3D` indicates that the implementation supports using a 2D view of a 3D image in a descriptor of type `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` if the image is created using `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT`.

## [](#_description)Description

If the `VkPhysicalDeviceImage2DViewOf3DFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceImage2DViewOf3DFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImage2DViewOf3DFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceImage2DViewOf3DFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_2D_VIEW_OF_3D_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_image\_2d\_view\_of\_3d](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_2d_view_of_3d.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImage2DViewOf3DFeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0