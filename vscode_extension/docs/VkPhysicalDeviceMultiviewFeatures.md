# VkPhysicalDeviceMultiviewFeatures(3) Manual Page

## Name

VkPhysicalDeviceMultiviewFeatures - Structure describing multiview features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMultiviewFeatures` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceMultiviewFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           multiview;
    VkBool32           multiviewGeometryShader;
    VkBool32           multiviewTessellationShader;
} VkPhysicalDeviceMultiviewFeatures;
```

or the equivalent

```c++
// Provided by VK_KHR_multiview
typedef VkPhysicalDeviceMultiviewFeatures VkPhysicalDeviceMultiviewFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`multiview` specifies whether the implementation supports multiview rendering within a render pass. If this feature is not enabled, the view mask of each subpass **must** always be zero.
- []()`multiviewGeometryShader` specifies whether the implementation supports multiview rendering within a render pass, with [geometry shaders](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#geometry). If this feature is not enabled, then a pipeline compiled against a subpass with a non-zero view mask **must** not include a geometry shader.
- []()`multiviewTessellationShader` specifies whether the implementation supports multiview rendering within a render pass, with [tessellation shaders](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#tessellation). If this feature is not enabled, then a pipeline compiled against a subpass with a non-zero view mask **must** not include any tessellation shaders.

If the `VkPhysicalDeviceMultiviewFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMultiviewFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage

- [](#VUID-VkPhysicalDeviceMultiviewFeatures-multiviewGeometryShader-00580)VUID-VkPhysicalDeviceMultiviewFeatures-multiviewGeometryShader-00580  
  If `multiviewGeometryShader` is enabled then `multiview` **must** also be enabled
- [](#VUID-VkPhysicalDeviceMultiviewFeatures-multiviewTessellationShader-00581)VUID-VkPhysicalDeviceMultiviewFeatures-multiviewTessellationShader-00581  
  If `multiviewTessellationShader` is enabled then `multiview` **must** also be enabled

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMultiviewFeatures-sType-sType)VUID-VkPhysicalDeviceMultiviewFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMultiviewFeatures).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0