# VkPhysicalDeviceMemoryProperties(3) Manual Page

## Name

VkPhysicalDeviceMemoryProperties - Structure specifying physical device memory properties



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMemoryProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPhysicalDeviceMemoryProperties {
    uint32_t        memoryTypeCount;
    VkMemoryType    memoryTypes[VK_MAX_MEMORY_TYPES];
    uint32_t        memoryHeapCount;
    VkMemoryHeap    memoryHeaps[VK_MAX_MEMORY_HEAPS];
} VkPhysicalDeviceMemoryProperties;
```

## [](#_members)Members

- `memoryTypeCount` is the number of valid elements in the `memoryTypes` array.
- `memoryTypes` is an array of `VK_MAX_MEMORY_TYPES` [VkMemoryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryType.html) structures describing the *memory types* that **can** be used to access memory allocated from the heaps specified by `memoryHeaps`.
- `memoryHeapCount` is the number of valid elements in the `memoryHeaps` array.
- `memoryHeaps` is an array of `VK_MAX_MEMORY_HEAPS` [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html) structures describing the *memory heaps* from which memory **can** be allocated.

## [](#_description)Description

The `VkPhysicalDeviceMemoryProperties` structure describes a number of *memory heaps* as well as a number of *memory types* that **can** be used to access memory allocated in those heaps. Each heap describes a memory resource of a particular size, and each memory type describes a set of memory properties (e.g. host cached vs. uncached) that **can** be used with a given memory heap. Allocations using a particular memory type will consume resources from the heap indicated by that memory type’s heap index. More than one memory type **may** share each heap, and the heaps and memory types provide a mechanism to advertise an accurate size of the physical memory resources while allowing the memory to be used with a variety of different properties.

The number of memory heaps is given by `memoryHeapCount` and is less than or equal to `VK_MAX_MEMORY_HEAPS`. Each heap is described by an element of the `memoryHeaps` array as a [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html) structure. The number of memory types available across all memory heaps is given by `memoryTypeCount` and is less than or equal to `VK_MAX_MEMORY_TYPES`. Each memory type is described by an element of the `memoryTypes` array as a [VkMemoryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryType.html) structure.

At least one heap **must** include `VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` in [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html)::`flags`. If there are multiple heaps that all have similar performance characteristics, they **may** all include `VK_MEMORY_HEAP_DEVICE_LOCAL_BIT`. In a unified memory architecture (UMA) system there is often only a single memory heap which is considered to be equally “local” to the host and to the device, and such an implementation **must** advertise the heap as device-local.

Memory contents within a tile memory heap, denoted by `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM`, are only visible across the command buffers executed in a single command buffer submission batch within a [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html) or [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2.html) call. If the [VkPhysicalDeviceTileMemoryHeapPropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTileMemoryHeapPropertiesQCOM.html)::`queueSubmitBoundary` property is set, the visibility is extended across all batches in the submit call. Memory contents are discarded and made undefined after the respective submission batch or submit call. Tile memory **may** have different performance characteristics than non tile memory. Tile memory **can** be used simultaneously by command buffers in other queues without invalidating each others contents. Collectively, these rules define the *tile memory scope*.

Note

Tile memory heaps work differently than most heaps as it is allowing addressing on device cache memory. Therefore, the heap’s address space is aliased across the different queues, with each queue retaining its individual copy of the heap. The implementation takes care of any required saving and restoring of the tile memory contents.

Effectively, this means that the same address in the heap in different queues have simultaneously different defined contents and the contents has a lifespan scoped to the submit or batch for that specific queues.

Each memory type returned by [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceMemoryProperties.html) **must** have its `propertyFlags` set to one of the following values:

- 0
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT`
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT`
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT`
- `VK_MEMORY_PROPERTY_PROTECTED_BIT`
- `VK_MEMORY_PROPERTY_PROTECTED_BIT` | `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT`
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` |  
  `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD`
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` |  
  `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` |  
  `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` |  
  `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` |  
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` |  
  `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` |  
  `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` |  
  `VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV`

There **must** be at least one memory type with both the `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` and `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` bits set in its `propertyFlags`. There **must** be at least one memory type with the `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` bit set in its `propertyFlags`. If the [`deviceCoherentMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceCoherentMemory) feature is enabled, there **must** be at least one memory type with the `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` bit set in its `propertyFlags`.

For each pair of elements **X** and **Y** returned in `memoryTypes`, **X** **must** be placed at a lower index position than **Y** if:

- the set of bit flags returned in the `propertyFlags` member of **X** is a strict subset of the set of bit flags returned in the `propertyFlags` member of **Y**; or
- the `propertyFlags` members of **X** and **Y** are equal, and **X** belongs to a memory heap with greater performance (as determined in an implementation-specific manner) ; or
- the `propertyFlags` members of **Y** includes `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` or `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD` and **X** does not

Note

There is no ordering requirement between **X** and **Y** elements for the case their `propertyFlags` members are not in a subset relation. That potentially allows more than one possible way to order the same set of memory types. Notice that the [list of all allowed memory property flag combinations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-device-bitmask-list) is written in a valid order. But if instead `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` was before `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` | `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT`, the list would still be in a valid order.

There may be a performance penalty for using device coherent or uncached device memory types, and using these accidentally is undesirable. In order to avoid this, memory types with these properties always appear at the end of the list; but are subject to the same rules otherwise.

This ordering requirement enables applications to use a simple search loop to select the desired memory type along the lines of:

```c++
// Find a memory in `memoryTypeBitsRequirement` that includes all of `requiredProperties`
int32_t findProperties(const VkPhysicalDeviceMemoryProperties* pMemoryProperties,
                       uint32_t memoryTypeBitsRequirement,
                       VkMemoryPropertyFlags requiredProperties) {
    const uint32_t memoryCount = pMemoryProperties->memoryTypeCount;
    for (uint32_t memoryIndex = 0; memoryIndex < memoryCount; ++memoryIndex) {
        const uint32_t memoryTypeBits = (1 << memoryIndex);
        const bool isRequiredMemoryType = memoryTypeBitsRequirement & memoryTypeBits;

        const VkMemoryPropertyFlags properties =
            pMemoryProperties->memoryTypes[memoryIndex].propertyFlags;
        const bool hasRequiredProperties =
            (properties & requiredProperties) == requiredProperties;

        if (isRequiredMemoryType && hasRequiredProperties)
            return static_cast<int32_t>(memoryIndex);
    }

    // failed to find memory type
    return -1;
}

// Try to find an optimal memory type, or if it does not exist try fallback memory type
// `device` is the VkDevice
// `image` is the VkImage that requires memory to be bound
// `memoryProperties` properties as returned by vkGetPhysicalDeviceMemoryProperties
// `requiredProperties` are the property flags that must be present
// `optimalProperties` are the property flags that are preferred by the application
VkMemoryRequirements memoryRequirements;
vkGetImageMemoryRequirements(device, image, &memoryRequirements);
int32_t memoryType =
    findProperties(&memoryProperties, memoryRequirements.memoryTypeBits, optimalProperties);
if (memoryType == -1) // not found; try fallback properties
    memoryType =
        findProperties(&memoryProperties, memoryRequirements.memoryTypeBits, requiredProperties);
```

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html), [VkMemoryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryType.html), [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties2.html), [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceMemoryProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMemoryProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0