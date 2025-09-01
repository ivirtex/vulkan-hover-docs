# VkPhysicalDeviceCudaKernelLaunchPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceCudaKernelLaunchPropertiesNV - Structure describing the compute capability version available



## [](#_c_specification)C Specification

The `VkPhysicalDeviceCudaKernelLaunchPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkPhysicalDeviceCudaKernelLaunchPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           computeCapabilityMinor;
    uint32_t           computeCapabilityMajor;
} VkPhysicalDeviceCudaKernelLaunchPropertiesNV;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceCudaKernelLaunchPropertiesNV` structure describe the following features:

## [](#_description)Description

- []()`computeCapabilityMinor` indicates the minor version number of the compute code.
- []()`computeCapabilityMajor` indicates the major version number of the compute code.

If the `VkPhysicalDeviceCudaKernelLaunchPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceCudaKernelLaunchPropertiesNV-sType-sType)VUID-VkPhysicalDeviceCudaKernelLaunchPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUDA_KERNEL_LAUNCH_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceCudaKernelLaunchPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0