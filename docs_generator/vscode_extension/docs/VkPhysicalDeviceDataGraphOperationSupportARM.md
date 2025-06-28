# VkPhysicalDeviceDataGraphOperationSupportARM(3) Manual Page

## Name

VkPhysicalDeviceDataGraphOperationSupportARM - Structure describing an operation or set of operations supported by a data graph processing engine



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDataGraphOperationSupportARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkPhysicalDeviceDataGraphOperationSupportARM {
    VkPhysicalDeviceDataGraphOperationTypeARM    operationType;
    char                                         name[VK_MAX_PHYSICAL_DEVICE_DATA_GRAPH_OPERATION_SET_NAME_SIZE_ARM];
    uint32_t                                     version;
} VkPhysicalDeviceDataGraphOperationSupportARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `operationType` is a [VkPhysicalDeviceDataGraphOperationTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationTypeARM.html) enum specifying the type of the operation whose support is being described.
- `name` is a pointer to a null-terminated UTF-8 string specifying the name of the operation whose support is being described.
- `version` is an integer specifying the version of the operation whose support is being described.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDataGraphOperationSupportARM-operationType-parameter)VUID-VkPhysicalDeviceDataGraphOperationSupportARM-operationType-parameter  
  `operationType` **must** be a valid [VkPhysicalDeviceDataGraphOperationTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationTypeARM.html) value
- [](#VUID-VkPhysicalDeviceDataGraphOperationSupportARM-name-parameter)VUID-VkPhysicalDeviceDataGraphOperationSupportARM-name-parameter  
  `name` **must** be a null-terminated UTF-8 string whose length is less than or equal to VK\_MAX\_PHYSICAL\_DEVICE\_DATA\_GRAPH\_OPERATION\_SET\_NAME\_SIZE\_ARM

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkPhysicalDeviceDataGraphOperationTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationTypeARM.html), [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDataGraphOperationSupportARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0