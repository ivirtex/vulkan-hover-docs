# VkBindVertexBufferIndirectCommandNV(3) Manual Page

## Name

VkBindVertexBufferIndirectCommandNV - Structure specifying input data
for a single vertex buffer command token



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindVertexBufferIndirectCommandNV` structure specifies the input
data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV` token.

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkBindVertexBufferIndirectCommandNV {
    VkDeviceAddress    bufferAddress;
    uint32_t           size;
    uint32_t           stride;
} VkBindVertexBufferIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `bufferAddress` specifies a physical address of the
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) used as vertex input binding.

- `size` is the byte size range which is available for this operation
  from the provided address.

- `stride` is the byte size stride for this vertex input binding as in
  `VkVertexInputBindingDescription`::`stride`. It is only used if
  `VkIndirectCommandsLayoutTokenNV`::`vertexDynamicStride` was set,
  otherwise the stride is inherited from the current bound graphics
  pipeline.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindVertexBufferIndirectCommandNV-None-02949"
  id="VUID-VkBindVertexBufferIndirectCommandNV-None-02949"></a>
  VUID-VkBindVertexBufferIndirectCommandNV-None-02949  
  The buffer’s usage flag from which the address was acquired **must**
  have the `VK_BUFFER_USAGE_VERTEX_BUFFER_BIT` bit set

- <a href="#VUID-VkBindVertexBufferIndirectCommandNV-None-02950"
  id="VUID-VkBindVertexBufferIndirectCommandNV-None-02950"></a>
  VUID-VkBindVertexBufferIndirectCommandNV-None-02950  
  Each element of the buffer from which the address was acquired and
  that is non-sparse **must** be bound completely and contiguously to a
  single `VkDeviceMemory` object

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindVertexBufferIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
