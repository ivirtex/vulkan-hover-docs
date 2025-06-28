# VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM - Structure describing multiview per view render areas features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_multiview_per_view_render_areas
typedef struct VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           multiviewPerViewRenderAreas;
} VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM;
```

## [](#_members)Members

This structure describes the following features:

- []()`multiviewPerViewRenderAreas` indicates that the implementation supports multiview per-view render areas.

## [](#_description)Description

If the `VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM-sType-sType)VUID-VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_RENDER_AREAS_FEATURES_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_multiview\_per\_view\_render\_areas](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_multiview_per_view_render_areas.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0