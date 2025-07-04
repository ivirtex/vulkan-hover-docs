# VkPhysicalDeviceVulkan11Properties(3) Manual Page

## Name

VkPhysicalDeviceVulkan11Properties - Structure specifying physical device properties for functionality promoted to Vulkan 1.1



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVulkan11Properties` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceVulkan11Properties {
    VkStructureType            sType;
    void*                      pNext;
    uint8_t                    deviceUUID[VK_UUID_SIZE];
    uint8_t                    driverUUID[VK_UUID_SIZE];
    uint8_t                    deviceLUID[VK_LUID_SIZE];
    uint32_t                   deviceNodeMask;
    VkBool32                   deviceLUIDValid;
    uint32_t                   subgroupSize;
    VkShaderStageFlags         subgroupSupportedStages;
    VkSubgroupFeatureFlags     subgroupSupportedOperations;
    VkBool32                   subgroupQuadOperationsInAllStages;
    VkPointClippingBehavior    pointClippingBehavior;
    uint32_t                   maxMultiviewViewCount;
    uint32_t                   maxMultiviewInstanceIndex;
    VkBool32                   protectedNoFault;
    uint32_t                   maxPerSetDescriptors;
    VkDeviceSize               maxMemoryAllocationSize;
} VkPhysicalDeviceVulkan11Properties;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `deviceUUID` is an array of `VK_UUID_SIZE` `uint8_t` values representing a universally unique identifier for the device.
- `driverUUID` is an array of `VK_UUID_SIZE` `uint8_t` values representing a universally unique identifier for the driver build in use by the device.
- `deviceLUID` is an array of `VK_LUID_SIZE` `uint8_t` values representing a locally unique identifier for the device.
- `deviceNodeMask` is a `uint32_t` bitfield identifying the node within a linked device adapter corresponding to the device.
- `deviceLUIDValid` is a boolean value that will be `VK_TRUE` if `deviceLUID` contains a valid LUID and `deviceNodeMask` contains a valid node mask, and `VK_FALSE` if they do not.

<!--THE END-->

- []()`subgroupSize` is the default number of invocations in each subgroup. `subgroupSize` is at least 1 if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`. `subgroupSize` is a power-of-two.
- []()`subgroupSupportedStages` is a bitfield of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) describing the shader stages that [group operations](#shaders-group-operations) with [subgroup scope](#shaders-scope-subgroup) are supported in. `subgroupSupportedStages` will have the `VK_SHADER_STAGE_COMPUTE_BIT` bit set if any of the physical device’s queues support `VK_QUEUE_COMPUTE_BIT`.
- []()`subgroupSupportedOperations` is a bitmask of [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlagBits.html) specifying the sets of [group operations](#shaders-group-operations) with [subgroup scope](#shaders-scope-subgroup) supported on this device. `subgroupSupportedOperations` will have the `VK_SUBGROUP_FEATURE_BASIC_BIT` bit set if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`.
- []()`subgroupQuadOperationsInAllStages` is a boolean specifying whether [quad group operations](#shaders-quad-operations) are available in all stages, or are restricted to fragment and compute stages.
- []()`pointClippingBehavior` is a [VkPointClippingBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPointClippingBehavior.html) value specifying the point clipping behavior supported by the implementation.
- []()`maxMultiviewViewCount` is one greater than the maximum view index that **can** be used in a subpass.
- []()`maxMultiviewInstanceIndex` is the maximum valid value of instance index allowed to be generated by a drawing command recorded within a subpass of a multiview render pass instance.
- []()`protectedNoFault` specifies how an implementation behaves when an application attempts to write to unprotected memory in a protected queue operation, read from protected memory in an unprotected queue operation, or perform a query in a protected queue operation. If this limit is `VK_TRUE`, such writes will be discarded or have undefined values written, reads and queries will return undefined values. If this limit is `VK_FALSE`, applications **must** not perform these operations. See [\[memory-protected-access-rules\]](#memory-protected-access-rules) for more information.
- []()`maxPerSetDescriptors` is a maximum number of descriptors (summed over all descriptor types) in a single descriptor set that is guaranteed to satisfy any implementation-dependent constraints on the size of a descriptor set itself. Applications **can** query whether a descriptor set that goes beyond this limit is supported using [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSupport.html).
- []()`maxMemoryAllocationSize` is the maximum size of a memory allocation that **can** be created, even if there is more space available in the heap. If [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html)::`allocationSize` is larger the error `VK_ERROR_OUT_OF_DEVICE_MEMORY` **may** be returned.

If the `VkPhysicalDeviceVulkan11Properties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These properties correspond to Vulkan 1.1 functionality.

The members of `VkPhysicalDeviceVulkan11Properties` have the same values as the corresponding members of [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html), [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html), [VkPhysicalDevicePointClippingProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePointClippingProperties.html), [VkPhysicalDeviceMultiviewProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultiviewProperties.html), [VkPhysicalDeviceProtectedMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProtectedMemoryProperties.html), and [VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance3Properties.html).

Note

The `subgroupSupportedStages`, `subgroupSupportedOperations`, and `subgroupQuadOperationsInAllStages` members of this structure correspond respectively to the [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedStages`, [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations`, and [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupProperties.html)::`quadOperationsInAllStages` members, but add the `subgroup` prefix to the member name.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVulkan11Properties-sType-sType)VUID-VkPhysicalDeviceVulkan11Properties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_1_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkPointClippingBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPointClippingBehavior.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubgroupFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubgroupFeatureFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVulkan11Properties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0