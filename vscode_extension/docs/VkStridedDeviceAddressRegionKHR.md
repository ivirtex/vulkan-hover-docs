# VkStridedDeviceAddressRegionKHR(3) Manual Page

## Name

VkStridedDeviceAddressRegionKHR - Structure specifying a region of device addresses with a stride



## [](#_c_specification)C Specification

The `VkStridedDeviceAddressRegionKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkStridedDeviceAddressRegionKHR {
    VkDeviceAddress    deviceAddress;
    VkDeviceSize       stride;
    VkDeviceSize       size;
} VkStridedDeviceAddressRegionKHR;
```

## [](#_members)Members

- `deviceAddress` is the device address (as returned by the [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html) command) at which the region starts, or zero if the region is unused.
- `stride` is the byte stride between consecutive elements.
- `size` is the size in bytes of the region starting at `deviceAddress`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkStridedDeviceAddressRegionKHR-size-04631)VUID-VkStridedDeviceAddressRegionKHR-size-04631  
  If `size` is not zero, all addresses between `deviceAddress` and `deviceAddress` + `size` - 1 **must** be in the buffer device address range of the same buffer
- [](#VUID-VkStridedDeviceAddressRegionKHR-size-04632)VUID-VkStridedDeviceAddressRegionKHR-size-04632  
  If `size` is not zero, `stride` **must** be less than or equal to the size of the buffer from which `deviceAddress` was queried

Valid Usage (Implicit)

- [](#VUID-VkStridedDeviceAddressRegionKHR-deviceAddress-parameter)VUID-VkStridedDeviceAddressRegionKHR-deviceAddress-parameter  
  If `deviceAddress` is not `0`, `deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysIndirectKHR.html), [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkStridedDeviceAddressRegionKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0