# VkPhysicalDeviceShaderEnqueuePropertiesAMDX(3) Manual Page

## Name

VkPhysicalDeviceShaderEnqueuePropertiesAMDX - Structure describing
shader enqueue limits of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderEnqueuePropertiesAMDX` structure is defined
as:

``` c
// Provided by VK_AMDX_shader_enqueue
typedef struct VkPhysicalDeviceShaderEnqueuePropertiesAMDX {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxExecutionGraphDepth;
    uint32_t           maxExecutionGraphShaderOutputNodes;
    uint32_t           maxExecutionGraphShaderPayloadSize;
    uint32_t           maxExecutionGraphShaderPayloadCount;
    uint32_t           executionGraphDispatchAddressAlignment;
} VkPhysicalDeviceShaderEnqueuePropertiesAMDX;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceShaderEnqueuePropertiesAMDX`
structure describe the following limits:

## <a href="#_description" class="anchor"></a>Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-maxExecutionGraphDepth"></span>
  `maxExecutionGraphDepth` defines the maximum node chain depth in the
  graph. The dispatched node is at depth 1 and the node enqueued by it
  is at depth 2, and so on. If a node enqueues itself, each recursive
  enqueue increases the depth by 1 as well.

- <span id="limits-maxExecutionGraphShaderOutputNodes"></span>
  `maxExecutionGraphShaderOutputNodes` specifies the maximum number of
  unique nodes that can be dispatched from a single shader, and must be
  at least 256.

- <span id="limits-maxExecutionGraphShaderPayloadSize"></span>
  `maxExecutionGraphShaderPayloadSize` specifies the maximum total size
  of payload declarations in a shader. For any payload declarations that
  share resources, indicated by `NodeSharesPayloadLimitsWithAMDX`
  decorations, the maximum size of each set of shared payload
  declarations is taken. The sum of each shared set’s maximum size and
  the size of each unshared payload is counted against this limit.

- <span id="limits-maxExecutionGraphShaderPayloadCount"></span>
  `maxExecutionGraphShaderPayloadCount` specifies the maximum number of
  output payloads that can be initialized in a single workgroup.

- <span id="limits-executionGraphDispatchAddressAlignment"></span>
  `executionGraphDispatchAddressAlignment` specifies the alignment of
  non-scratch [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) arguments consumed
  by graph dispatch commands.

If the `VkPhysicalDeviceShaderEnqueuePropertiesAMDX` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderEnqueuePropertiesAMDX-sType-sType"
  id="VUID-VkPhysicalDeviceShaderEnqueuePropertiesAMDX-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderEnqueuePropertiesAMDX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ENQUEUE_PROPERTIES_AMDX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderEnqueuePropertiesAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
