# VkDispatchGraphInfoAMDX(3) Manual Page

## Name

VkDispatchGraphInfoAMDX - Structure specifying node parameters for
execution graph dispatch



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDispatchGraphInfoAMDX` structure is defined as:

``` c
// Provided by VK_AMDX_shader_enqueue
typedef struct VkDispatchGraphInfoAMDX {
    uint32_t                          nodeIndex;
    uint32_t                          payloadCount;
    VkDeviceOrHostAddressConstAMDX    payloads;
    uint64_t                          payloadStride;
} VkDispatchGraphInfoAMDX;
```

## <a href="#_members" class="anchor"></a>Members

- `nodeIndex` is the index of a node in an execution graph to be
  dispatched.

- `payloadCount` is the number of payloads to dispatch for the specified
  node.

- `payloads` is a device or host address pointer to a flat array of
  payloads with size equal to the product of `payloadCount` and
  `payloadStride`

- `payloadStride` is the byte stride between successive payloads in
  `payloads`

## <a href="#_description" class="anchor"></a>Description

Whether `payloads` is consumed as a device or host pointer is defined by
the command this structure is used in.

Valid Usage

- <a href="#VUID-VkDispatchGraphInfoAMDX-payloadCount-09171"
  id="VUID-VkDispatchGraphInfoAMDX-payloadCount-09171"></a>
  VUID-VkDispatchGraphInfoAMDX-payloadCount-09171  
  `payloadCount` **must** be no greater than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxExecutionGraphShaderPayloadCount"
  target="_blank"
  rel="noopener"><code>maxExecutionGraphShaderPayloadCount</code></a>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkDeviceOrHostAddressConstAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstAMDX.html),
[VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDispatchGraphCountInfoAMDX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDispatchGraphInfoAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
