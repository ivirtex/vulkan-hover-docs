# VkPhysicalDeviceCudaKernelLaunchFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceCudaKernelLaunchFeaturesNV - Structure describing
whether cuda kernel launch is supported by the implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCudaKernelLaunchFeaturesNV` structure is defined
as:

``` c
// Provided by VK_NV_cuda_kernel_launch
typedef struct VkPhysicalDeviceCudaKernelLaunchFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           cudaKernelLaunchFeatures;
} VkPhysicalDeviceCudaKernelLaunchFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- <span id="features-cudaKernelLaunchFeatures"></span>
  `cudaKernelLaunchFeatures` is non-zero if cuda kernel launch is
  supported.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceCudaKernelLaunchFeaturesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceCudaKernelLaunchFeaturesNV` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceCudaKernelLaunchFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceCudaKernelLaunchFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceCudaKernelLaunchFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUDA_KERNEL_LAUNCH_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCudaKernelLaunchFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
