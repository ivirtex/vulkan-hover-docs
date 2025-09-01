# VkPhysicalDeviceHostImageCopyFeatures(3) Manual Page

## Name

VkPhysicalDeviceHostImageCopyFeatures - Structure indicating support for copies to or from images from host memory



## [](#_c_specification)C Specification

The `VkPhysicalDeviceHostImageCopyFeatures` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceHostImageCopyFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           hostImageCopy;
} VkPhysicalDeviceHostImageCopyFeatures;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkPhysicalDeviceHostImageCopyFeatures VkPhysicalDeviceHostImageCopyFeaturesEXT;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`hostImageCopy` indicates that the implementation supports copying from host memory to images using the [vkCopyMemoryToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToImage.html) command, copying from images to host memory using the [vkCopyImageToMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToMemory.html) command, and copying between images using the [vkCopyImageToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToImage.html) command.

If the `VkPhysicalDeviceHostImageCopyFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceHostImageCopyFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceHostImageCopyFeatures-sType-sType)VUID-VkPhysicalDeviceHostImageCopyFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_IMAGE_COPY_FEATURES`

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceHostImageCopyFeatures).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0