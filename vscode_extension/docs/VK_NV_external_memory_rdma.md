# VK\_NV\_external\_memory\_rdma(3) Manual Page

## Name

VK\_NV\_external\_memory\_rdma - device extension



## [](#_registered_extension_number)Registered Extension Number

372

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Carsten Rohde [\[GitHub\]crohde](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_external_memory_rdma%5D%20%40crohde%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_external_memory_rdma%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-04-19

**IP Status**

No known IP claims.

**Contributors**

- Carsten Rohde, NVIDIA

## [](#_description)Description

This extension adds support for allocating memory which can be used for remote direct memory access (RDMA) from other devices.

## [](#_new_base_types)New Base Types

- [VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRemoteAddressNV.html)

## [](#_new_commands)New Commands

- [vkGetMemoryRemoteAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryRemoteAddressNV.html)

## [](#_new_structures)New Structures

- [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetRemoteAddressInfoNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExternalMemoryRDMAFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalMemoryRDMAFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_EXTERNAL_MEMORY_RDMA_EXTENSION_NAME`
- `VK_NV_EXTERNAL_MEMORY_RDMA_SPEC_VERSION`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_RDMA_ADDRESS_BIT_NV`
- Extending [VkMemoryPropertyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryPropertyFlagBits.html):
  
  - `VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MEMORY_GET_REMOTE_ADDRESS_INFO_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_RDMA_FEATURES_NV`

## [](#_issues)Issues

## [](#_examples)Examples

```cpp
VkPhysicalDeviceMemoryBudgetPropertiesEXT memoryBudgetProperties = { VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT };
VkPhysicalDeviceMemoryProperties2 memoryProperties2 = { VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2, &memoryBudgetProperties };
vkGetPhysicalDeviceMemoryProperties2(physicalDevice, &memoryProperties2);
uint32_t heapIndex = (uint32_t)-1;
for (uint32_t memoryType = 0; memoryType < memoryProperties2.memoryProperties.memoryTypeCount; memoryType++) {
    if (memoryProperties2.memoryProperties.memoryTypes[memoryType].propertyFlags & VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV) {
        heapIndex = memoryProperties2.memoryProperties.memoryTypes[memoryType].heapIndex;
        break;
    }
}
if ((heapIndex == (uint32_t)-1) ||
    (memoryBudgetProperties.heapBudget[heapIndex] < size)) {
    return;
}

VkPhysicalDeviceExternalBufferInfo externalBufferInfo = { VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO };
externalBufferInfo.usage = VK_BUFFER_USAGE_TRANSFER_SRC_BIT | VK_BUFFER_USAGE_TRANSFER_DST_BIT;
externalBufferInfo.handleType = VK_EXTERNAL_MEMORY_HANDLE_TYPE_RDMA_ADDRESS_BIT_NV;

VkExternalBufferProperties externalBufferProperties = { VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES };
vkGetPhysicalDeviceExternalBufferProperties(physicalDevice, &externalBufferInfo, &externalBufferProperties);

if (!(externalBufferProperties.externalMemoryProperties.externalMemoryFeatures & VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT)) {
    return;
}

VkExternalMemoryBufferCreateInfo externalMemoryBufferCreateInfo = { VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO };
externalMemoryBufferCreateInfo.handleTypes = VK_EXTERNAL_MEMORY_HANDLE_TYPE_RDMA_ADDRESS_BIT_NV;

VkBufferCreateInfo bufferCreateInfo = { VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO, &externalMemoryBufferCreateInfo };
bufferCreateInfo.size = size;
bufferCreateInfo.usage = VK_BUFFER_USAGE_TRANSFER_SRC_BIT | VK_BUFFER_USAGE_TRANSFER_DST_BIT;

VkMemoryRequirements mem_reqs;
vkCreateBuffer(device, &bufferCreateInfo, NULL, &buffer);
vkGetBufferMemoryRequirements(device, buffer, &mem_reqs);

VkExportMemoryAllocateInfo exportMemoryAllocateInfo = { VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO };
exportMemoryAllocateInfo.handleTypes = VK_EXTERNAL_MEMORY_HANDLE_TYPE_RDMA_ADDRESS_BIT_NV;

// Find memory type index
uint32_t i = 0;
for (; i < VK_MAX_MEMORY_TYPES; i++) {
    if ((mem_reqs.memoryTypeBits & (1 << i)) &&
        (memoryProperties.memoryTypes[i].propertyFlags & VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV)) {
        break;
    }
}

VkMemoryAllocateInfo memAllocInfo = { VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO, &exportMemoryAllocateInfo };
memAllocInfo.allocationSize = mem_reqs.size;
memAllocInfo.memoryTypeIndex = i;

vkAllocateMemory(device, &memAllocInfo, NULL, &mem);
vkBindBufferMemory(device, buffer, mem, 0);

VkMemoryGetRemoteAddressInfoNV getMemoryRemoteAddressInfo = { VK_STRUCTURE_TYPE_MEMORY_GET_REMOTE_ADDRESS_INFO_NV };
getMemoryRemoteAddressInfo.memory = mem;
getMemoryRemoteAddressInfo.handleType = VK_EXTERNAL_MEMORY_HANDLE_TYPE_RDMA_ADDRESS_BIT_NV;

VkRemoteAddressNV rdmaAddress;
vkGetMemoryRemoteAddressNV(device, &getMemoryRemoteAddressInfo, &rdmaAddress);
// address returned in 'rdmaAddress' can be used by external devices to initiate RDMA transfers
```

## [](#_version_history)Version History

- Revision 1, 2020-12-15 (Carsten Rohde)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_external_memory_rdma).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0