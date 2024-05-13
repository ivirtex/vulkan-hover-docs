# VK_NV_external_memory_rdma(3) Manual Page

## Name

VK_NV_external_memory_rdma - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

372

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Carsten Rohde <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_external_memory_rdma%5D%20@crohde%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_external_memory_rdma%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>crohde</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-04-19

**IP Status**  
No known IP claims.

**Contributors**  
- Carsten Rohde, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for allocating memory which can be used for
remote direct memory access (RDMA) from other devices.

## <a href="#_new_base_types" class="anchor"></a>New Base Types

- [VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRemoteAddressNV.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetMemoryRemoteAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryRemoteAddressNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetRemoteAddressInfoNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceExternalMemoryRDMAFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalMemoryRDMAFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_EXTERNAL_MEMORY_RDMA_EXTENSION_NAME`

- `VK_NV_EXTERNAL_MEMORY_RDMA_SPEC_VERSION`

- Extending
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html):

  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_RDMA_ADDRESS_BIT_NV`

- Extending [VkMemoryPropertyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlagBits.html):

  - `VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MEMORY_GET_REMOTE_ADDRESS_INFO_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_RDMA_FEATURES_NV`

## <a href="#_issues" class="anchor"></a>Issues

## <a href="#_examples" class="anchor"></a>Examples

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-12-15 (Carsten Rohde)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_external_memory_rdma"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
