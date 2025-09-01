# VkDispatchGraphInfoAMDX(3) Manual Page

## Name

VkDispatchGraphInfoAMDX - Structure specifying node parameters for execution graph dispatch



## [](#_c_specification)C Specification

The `VkDispatchGraphInfoAMDX` structure is defined as:

```c++
// Provided by VK_AMDX_shader_enqueue
typedef struct VkDispatchGraphInfoAMDX {
    uint32_t                          nodeIndex;
    uint32_t                          payloadCount;
    VkDeviceOrHostAddressConstAMDX    payloads;
    uint64_t                          payloadStride;
} VkDispatchGraphInfoAMDX;
```

## [](#_members)Members

- `nodeIndex` is the index of a node in an execution graph to be dispatched.
- `payloadCount` is the number of payloads to dispatch for the specified node.
- `payloads` is a device or host address pointer to a flat array of payloads with size equal to the product of `payloadCount` and `payloadStride`
- `payloadStride` is the byte stride between successive payloads in `payloads`

## [](#_description)Description

Whether `payloads` is consumed as a device or host pointer is defined by the command this structure is used in.

Valid Usage

- [](#VUID-VkDispatchGraphInfoAMDX-payloadCount-09171)VUID-VkDispatchGraphInfoAMDX-payloadCount-09171  
  `payloadCount` **must** be no greater than [`maxExecutionGraphShaderPayloadCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxExecutionGraphShaderPayloadCount)

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkDeviceOrHostAddressConstAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstAMDX.html), [VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphCountInfoAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDispatchGraphInfoAMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0