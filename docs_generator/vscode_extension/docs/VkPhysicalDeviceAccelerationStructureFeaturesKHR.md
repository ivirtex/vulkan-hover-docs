# VkPhysicalDeviceAccelerationStructureFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceAccelerationStructureFeaturesKHR - Structure describing the acceleration structure features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceAccelerationStructureFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkPhysicalDeviceAccelerationStructureFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           accelerationStructure;
    VkBool32           accelerationStructureCaptureReplay;
    VkBool32           accelerationStructureIndirectBuild;
    VkBool32           accelerationStructureHostCommands;
    VkBool32           descriptorBindingAccelerationStructureUpdateAfterBind;
} VkPhysicalDeviceAccelerationStructureFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`accelerationStructure` indicates whether the implementation supports the acceleration structure functionality. See [Acceleration Structures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure).
- []()`accelerationStructureCaptureReplay` indicates whether the implementation supports saving and reusing acceleration structure device addresses, e.g. for trace capture and replay.
- []()`accelerationStructureIndirectBuild` indicates whether the implementation supports indirect acceleration structure build commands, e.g. [vkCmdBuildAccelerationStructuresIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructuresIndirectKHR.html).
- []()`accelerationStructureHostCommands` indicates whether the implementation supports host side acceleration structure commands, e.g. [vkBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBuildAccelerationStructuresKHR.html), [vkCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureKHR.html), [vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureToMemoryKHR.html), [vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToAccelerationStructureKHR.html), [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteAccelerationStructuresPropertiesKHR.html).
- []()`descriptorBindingAccelerationStructureUpdateAfterBind` indicates whether the implementation supports updating acceleration structure descriptors after a set is bound. If this feature is not enabled, `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used with `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`.

## [](#_description)Description

If the `VkPhysicalDeviceAccelerationStructureFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceAccelerationStructureFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceAccelerationStructureFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceAccelerationStructureFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceAccelerationStructureFeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0