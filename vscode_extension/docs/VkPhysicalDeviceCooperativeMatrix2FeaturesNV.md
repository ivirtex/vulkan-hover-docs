# VkPhysicalDeviceCooperativeMatrix2FeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceCooperativeMatrix2FeaturesNV - Structure describing cooperative matrix features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceCooperativeMatrix2FeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_cooperative_matrix2
typedef struct VkPhysicalDeviceCooperativeMatrix2FeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           cooperativeMatrixWorkgroupScope;
    VkBool32           cooperativeMatrixFlexibleDimensions;
    VkBool32           cooperativeMatrixReductions;
    VkBool32           cooperativeMatrixConversions;
    VkBool32           cooperativeMatrixPerElementOperations;
    VkBool32           cooperativeMatrixTensorAddressing;
    VkBool32           cooperativeMatrixBlockLoads;
} VkPhysicalDeviceCooperativeMatrix2FeaturesNV;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`cooperativeMatrixWorkgroupScope` indicates that the implementation supports workgroup scope cooperative matrices.
- []()`cooperativeMatrixFlexibleDimensions` indicates that the implementation supports cooperative matrix sizes that are a multiple of the granularity advertised in [VkCooperativeMatrixFlexibleDimensionsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeMatrixFlexibleDimensionsPropertiesNV.html).
- []()`cooperativeMatrixReductions` indicates that the implementation supports the `CooperativeMatrixReductionsNV` SPIR-V capability. This allows performing (row, column, 2x2, or all element) reductions on matrices.
- []()`cooperativeMatrixConversions` indicates that the implementation supports the `CooperativeMatrixConversionsNV` SPIR-V capability. This allows converting accumulator matrices to A or B matrices.
- []()`cooperativeMatrixPerElementOperations` indicates that the implementation supports the `CooperativeMatrixPerElementOperationsNV` SPIR-V capability. This allows performing element-wise operations on matrix elements using a callback function.
- []()`cooperativeMatrixTensorAddressing` indicates that the implementation supports the `TensorAddressingNV` and `CooperativeMatrixTensorAddressingNV` SPIR-V capabilities. This allows using tensor layout and tensor view types for matrix loads and stores.
- []()`cooperativeMatrixBlockLoads` indicates that the implementation supports the `CooperativeMatrixBlockLoadsNV` SPIR-V capability. This allows setting block size for loads and using a callback function to decode block elements.

## [](#_description)Description

If the `VkPhysicalDeviceCooperativeMatrix2FeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceCooperativeMatrix2FeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceCooperativeMatrix2FeaturesNV-sType-sType)VUID-VkPhysicalDeviceCooperativeMatrix2FeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_2_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_cooperative\_matrix2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_matrix2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceCooperativeMatrix2FeaturesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0