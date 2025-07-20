# vkCreateAccelerationStructureKHR(3) Manual Page

## Name

vkCreateAccelerationStructureKHR - Create a new acceleration structure object



## [](#_c_specification)C Specification

To create an acceleration structure, call:

```c++
// Provided by VK_KHR_acceleration_structure
VkResult vkCreateAccelerationStructureKHR(
    VkDevice                                    device,
    const VkAccelerationStructureCreateInfoKHR* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkAccelerationStructureKHR*                 pAccelerationStructure);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the acceleration structure object.
- `pCreateInfo` is a pointer to a [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html) structure containing parameters affecting creation of the acceleration structure.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pAccelerationStructure` is a pointer to a `VkAccelerationStructureKHR` handle in which the resulting acceleration structure object is returned.

## [](#_description)Description

Similar to other objects in Vulkan, the acceleration structure creation merely creates an object with a specific “shape”. The type and quantity of geometry that can be built into an acceleration structure is determined by the parameters of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html).

The acceleration structure data is stored in the object referred to by `VkAccelerationStructureCreateInfoKHR`::`buffer`. Once memory has been bound to that buffer, it **must** be populated by acceleration structure build or acceleration structure copy commands such as [vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructuresKHR.html), [vkBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBuildAccelerationStructuresKHR.html), [vkCmdCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureKHR.html), and [vkCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureKHR.html).

Note

The expected usage for a trace capture/replay tool is that it will serialize and later deserialize the acceleration structure data using acceleration structure copy commands. During capture the tool will use [vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureToMemoryKHR.html) or [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html) with a `mode` of `VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR`, and [vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToAccelerationStructureKHR.html) or [vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html) with a `mode` of `VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR` during replay.

Note

Memory does not need to be bound to the underlying buffer when [vkCreateAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateAccelerationStructureKHR.html) is called.

The input buffers passed to acceleration structure build commands will be referenced by the implementation for the duration of the command. After the command completes, the acceleration structure **may** hold a reference to any acceleration structure specified by an active instance contained therein. Apart from this referencing, acceleration structures **must** be fully self-contained. The application **can** reuse or free any memory which was used by the command as an input or as scratch without affecting the results of ray traversal.

Valid Usage

- [](#VUID-vkCreateAccelerationStructureKHR-accelerationStructure-03611)VUID-vkCreateAccelerationStructureKHR-accelerationStructure-03611  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkCreateAccelerationStructureKHR-deviceAddress-03488)VUID-vkCreateAccelerationStructureKHR-deviceAddress-03488  
  If [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`deviceAddress` is not zero, the [`accelerationStructureCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructureCaptureReplay) feature **must** be enabled
- [](#VUID-vkCreateAccelerationStructureKHR-device-03489)VUID-vkCreateAccelerationStructureKHR-device-03489  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreateAccelerationStructureKHR-device-parameter)VUID-vkCreateAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateAccelerationStructureKHR-pCreateInfo-parameter)VUID-vkCreateAccelerationStructureKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html) structure
- [](#VUID-vkCreateAccelerationStructureKHR-pAllocator-parameter)VUID-vkCreateAccelerationStructureKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateAccelerationStructureKHR-pAccelerationStructure-parameter)VUID-vkCreateAccelerationStructureKHR-pAccelerationStructure-parameter  
  `pAccelerationStructure` **must** be a valid pointer to a [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkCreateAccelerationStructureKHR-device-queuecount)VUID-vkCreateAccelerationStructureKHR-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateAccelerationStructureKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0