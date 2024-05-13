# VkPhysicalDeviceCudaKernelLaunchPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceCudaKernelLaunchPropertiesNV - Structure describing the
compute capability version available



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCudaKernelLaunchPropertiesNV` structure is defined
as:

``` c
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkPhysicalDeviceCudaKernelLaunchPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           computeCapabilityMinor;
    uint32_t           computeCapabilityMajor;
} VkPhysicalDeviceCudaKernelLaunchPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceCudaKernelLaunchPropertiesNV`
structure describe the following features:

## <a href="#_description" class="anchor"></a>Description

- <span id="limits-computeCapabilityMinor"></span>
  `computeCapabilityMinor` indicates the minor version number of the
  compute code.

- <span id="limits-computeCapabilityMajor"></span>
  `computeCapabilityMajor` indicates the major version number of the
  compute code.

If the `VkPhysicalDeviceCudaKernelLaunchPropertiesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceCudaKernelLaunchPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceCudaKernelLaunchPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceCudaKernelLaunchPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUDA_KERNEL_LAUNCH_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCudaKernelLaunchPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
