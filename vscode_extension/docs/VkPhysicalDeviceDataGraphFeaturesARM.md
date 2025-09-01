# VkPhysicalDeviceDataGraphFeaturesARM(3) Manual Page

## Name

VkPhysicalDeviceDataGraphFeaturesARM - Structure describing features to control data graph pipelines



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDataGraphFeaturesARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkPhysicalDeviceDataGraphFeaturesARM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           dataGraph;
    VkBool32           dataGraphUpdateAfterBind;
    VkBool32           dataGraphSpecializationConstants;
    VkBool32           dataGraphDescriptorBuffer;
    VkBool32           dataGraphShaderModule;
} VkPhysicalDeviceDataGraphFeaturesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`dataGraph` specifies whether data graph pipelines **can** be used.
- []()`dataGraphUpdateAfterBind` specifies whether data graph pipelines **can** be created with a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) that uses one or more [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) objects created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`dataGraphSpecializationConstants` specifies whether data graph pipelines **can** be created from shader modules that use specialization constants.
- []()`dataGraphDescriptorBuffer` specifies whether data graph pipelines **can** use descriptor buffers.
- []()`dataGraphShaderModule` specifies whether data graph pipelines **can** be created from a shader module.

## [](#_description)Description

If the `VkPhysicalDeviceDataGraphFeaturesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceDataGraphFeaturesARM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDataGraphFeaturesARM-sType-sType)VUID-VkPhysicalDeviceDataGraphFeaturesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DATA_GRAPH_FEATURES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDataGraphFeaturesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0