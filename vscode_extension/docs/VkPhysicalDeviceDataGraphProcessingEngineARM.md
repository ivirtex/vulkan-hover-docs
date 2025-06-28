# VkPhysicalDeviceDataGraphProcessingEngineARM(3) Manual Page

## Name

VkPhysicalDeviceDataGraphProcessingEngineARM - Structure describing a data graph processing engine supported by a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDataGraphProcessingEngineARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkPhysicalDeviceDataGraphProcessingEngineARM {
    VkPhysicalDeviceDataGraphProcessingEngineTypeARM    type;
    VkBool32                                            isForeign;
} VkPhysicalDeviceDataGraphProcessingEngineARM;
```

## [](#_members)Members

- `type` is a [VkPhysicalDeviceDataGraphProcessingEngineTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineTypeARM.html) that specifies the type of the processing engine.
- `isForeign` specifies whether the processing engine is foreign.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDataGraphProcessingEngineARM-type-parameter)VUID-VkPhysicalDeviceDataGraphProcessingEngineARM-type-parameter  
  `type` **must** be a valid [VkPhysicalDeviceDataGraphProcessingEngineTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineTypeARM.html) value

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html), [VkPhysicalDeviceDataGraphProcessingEngineTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineTypeARM.html), [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDataGraphProcessingEngineARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0