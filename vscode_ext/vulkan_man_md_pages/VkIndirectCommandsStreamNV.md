# VkIndirectCommandsStreamNV(3) Manual Page

## Name

VkIndirectCommandsStreamNV - Structure specifying input streams for
generated command tokens



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkIndirectCommandsStreamNV` structure specifies the input data for
one or more tokens at processing time.

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkIndirectCommandsStreamNV {
    VkBuffer        buffer;
    VkDeviceSize    offset;
} VkIndirectCommandsStreamNV;
```

## <a href="#_members" class="anchor"></a>Members

- `buffer` specifies the [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) storing the
  functional arguments for each sequence. These arguments **can** be
  written by the device.

- `offset` specified an offset into `buffer` where the arguments start.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkIndirectCommandsStreamNV-buffer-02942"
  id="VUID-VkIndirectCommandsStreamNV-buffer-02942"></a>
  VUID-VkIndirectCommandsStreamNV-buffer-02942  
  The `buffer`’s usage flag **must** have the
  `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set

- <a href="#VUID-VkIndirectCommandsStreamNV-offset-02943"
  id="VUID-VkIndirectCommandsStreamNV-offset-02943"></a>
  VUID-VkIndirectCommandsStreamNV-offset-02943  
  The `offset` **must** be aligned to
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`minIndirectCommandsBufferOffsetAlignment`

- <a href="#VUID-VkIndirectCommandsStreamNV-buffer-02975"
  id="VUID-VkIndirectCommandsStreamNV-buffer-02975"></a>
  VUID-VkIndirectCommandsStreamNV-buffer-02975  
  If `buffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

Valid Usage (Implicit)

- <a href="#VUID-VkIndirectCommandsStreamNV-buffer-parameter"
  id="VUID-VkIndirectCommandsStreamNV-buffer-parameter"></a>
  VUID-VkIndirectCommandsStreamNV-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIndirectCommandsStreamNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
