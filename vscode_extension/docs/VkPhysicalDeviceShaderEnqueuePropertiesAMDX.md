# VkPhysicalDeviceShaderEnqueuePropertiesAMDX(3) Manual Page

## Name

VkPhysicalDeviceShaderEnqueuePropertiesAMDX - Structure describing shader enqueue limits of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderEnqueuePropertiesAMDX` structure is defined as:

```c++
// Provided by VK_AMDX_shader_enqueue
typedef struct VkPhysicalDeviceShaderEnqueuePropertiesAMDX {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxExecutionGraphDepth;
    uint32_t           maxExecutionGraphShaderOutputNodes;
    uint32_t           maxExecutionGraphShaderPayloadSize;
    uint32_t           maxExecutionGraphShaderPayloadCount;
    uint32_t           executionGraphDispatchAddressAlignment;
    uint32_t           maxExecutionGraphWorkgroupCount[3];
    uint32_t           maxExecutionGraphWorkgroups;
} VkPhysicalDeviceShaderEnqueuePropertiesAMDX;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceShaderEnqueuePropertiesAMDX` structure describe the following limits:

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxExecutionGraphDepth` defines the maximum node chain depth in the graph. The dispatched node is at depth 1 and the node enqueued by it is at depth 2, and so on. If a node enqueues itself, each recursive enqueue increases the depth by 1 as well.
- []()`maxExecutionGraphShaderOutputNodes` specifies the maximum number of unique nodes that can be dispatched from a single shader, and **must** be at least 256.
- []()`maxExecutionGraphShaderPayloadSize` specifies the maximum total size of payload declarations in a shader. For any payload declarations that share resources, indicated by `NodeSharesPayloadLimitsWithAMDX` decorations, the maximum size of each set of shared payload declarations is taken. The sum of each shared set’s maximum size and the size of each unshared payload is counted against this limit.
- []()`maxExecutionGraphShaderPayloadCount` specifies the maximum number of output payloads that can be initialized in a single workgroup.
- []()`executionGraphDispatchAddressAlignment` specifies the alignment of non-scratch [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) arguments consumed by graph dispatch commands.
- []()`maxExecutionGraphWorkgroupCount`\[3] is the maximum number of local workgroups that a shader **can** be dispatched with in X, Y, and Z dimensions, respectively.
- []()`maxExecutionGraphWorkgroups` is the total number of local workgroups that a shader **can** be dispatched with.

If the `VkPhysicalDeviceShaderEnqueuePropertiesAMDX` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderEnqueuePropertiesAMDX-sType-sType)VUID-VkPhysicalDeviceShaderEnqueuePropertiesAMDX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ENQUEUE_PROPERTIES_AMDX`

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderEnqueuePropertiesAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0