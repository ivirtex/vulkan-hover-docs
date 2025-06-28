# VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV - Structure describing the device-generated compute features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           deviceGeneratedCompute;
    VkBool32           deviceGeneratedComputePipelines;
    VkBool32           deviceGeneratedComputeCaptureReplay;
} VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`deviceGeneratedCompute` indicates whether the implementation supports functionality to generate dispatch commands and push constants for the compute pipeline on the device. See [Device-Generated Commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#device-generated-commands).
- []()`deviceGeneratedComputePipelines` indicates whether the implementation supports functionality to generate commands to bind compute pipelines on the device. See [Device-Generated Commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#device-generated-commands).
- []()`deviceGeneratedComputeCaptureReplay` indicates whether the implementation supports functionality to capture compute pipeline address and reuse later for replay in [Device-Generated Commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#device-generated-commands).

## [](#_description)Description

If the `VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV-sType-sType)VUID-VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_COMPUTE_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands\_compute](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands_compute.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0