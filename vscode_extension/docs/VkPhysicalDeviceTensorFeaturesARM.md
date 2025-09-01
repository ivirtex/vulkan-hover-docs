# VkPhysicalDeviceTensorFeaturesARM(3) Manual Page

## Name

VkPhysicalDeviceTensorFeaturesARM - Structure describing tensor features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceTensorFeaturesARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkPhysicalDeviceTensorFeaturesARM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           tensorNonPacked;
    VkBool32           shaderTensorAccess;
    VkBool32           shaderStorageTensorArrayDynamicIndexing;
    VkBool32           shaderStorageTensorArrayNonUniformIndexing;
    VkBool32           descriptorBindingStorageTensorUpdateAfterBind;
    VkBool32           tensors;
} VkPhysicalDeviceTensorFeaturesARM;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceTensorFeaturesARM` structure describe the following features:

## [](#_description)Description

- []()`tensorNonPacked` indicates whether the implementation supports the creation of tensors that are not packed tensors.
- []()`shaderTensorAccess` indicates whether shader modules **can** declare the `TensorsARM` capability.
- []()`shaderStorageBufferArrayDynamicIndexing` indicates whether arrays of storage tensors **can** be indexed by dynamically uniform integer expressions in shader code. If this feature is not enabled, resources with a descriptor type of `VK_DESCRIPTOR_TYPE_TENSOR_ARM` **must** be indexed only by constant integral expressions when aggregated into arrays in shader code. This also indicates whether shader modules **can** declare the `StorageTensorArrayDynamicIndexingARM` capability.
- []()`shaderStorageTensorArrayNonUniformIndexing` indicates whether arrays of storage tensors **can** be indexed by non-uniform integer expressions in shader code. If this feature is not enabled, resources with a descriptor type of `VK_DESCRIPTOR_TYPE_TENSOR_ARM` **must** not be indexed by non-uniform integer expressions when aggregated into arrays in shader code. This also indicates whether shader modules **can** declare the `StorageTensorArrayNonUniformIndexingARM` capability.
- []()`descriptorBindingStorageTensorUpdateAfterBind` indicates whether the implementation supports updating storage tensor descriptors after a set is bound. If this feature is not enabled, `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used with `VK_DESCRIPTOR_TYPE_TENSOR_ARM`.
- []()`tensors` indicates whether the implementation supports tensor resources.

If the `VkPhysicalDeviceTensorFeaturesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceTensorFeaturesARM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceTensorFeaturesARM-sType-sType)VUID-VkPhysicalDeviceTensorFeaturesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TENSOR_FEATURES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceTensorFeaturesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0