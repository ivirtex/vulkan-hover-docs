# VkDescriptorBufferInfo(3) Manual Page

## Name

VkDescriptorBufferInfo - Structure specifying descriptor buffer
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDescriptorBufferInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorBufferInfo {
    VkBuffer        buffer;
    VkDeviceSize    offset;
    VkDeviceSize    range;
} VkDescriptorBufferInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `buffer` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or the buffer
  resource.

- `offset` is the offset in bytes from the start of `buffer`. Access to
  buffer memory via this descriptor uses addressing that is relative to
  this starting offset.

- `range` is the size in bytes that is used for this descriptor update,
  or `VK_WHOLE_SIZE` to use the range from `offset` to the end of the
  buffer.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>When setting <code>range</code> to <code>VK_WHOLE_SIZE</code>, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#buffer-info-effective-range"
  target="_blank" rel="noopener">effective range</a> <strong>must</strong>
  not be larger than the maximum range for the descriptor type (<a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxUniformBufferRange"
  target="_blank" rel="noopener"><code>maxUniformBufferRange</code></a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxStorageBufferRange"
  target="_blank" rel="noopener"><code>maxStorageBufferRange</code></a>).
  This means that <code>VK_WHOLE_SIZE</code> is not typically useful in
  the common case where uniform buffer descriptors are suballocated from a
  buffer that is much larger than
  <code>maxUniformBufferRange</code>.</p></td>
  </tr>
  </tbody>
  </table>

## <a href="#_description" class="anchor"></a>Description

For `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` and
`VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` descriptor types, `offset`
is the base offset from which the dynamic offset is applied and `range`
is the static size used for all dynamic offsets.

When `range` is `VK_WHOLE_SIZE` the effective range is calculated at
[vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html) is by taking the
size of `buffer` minus the `offset`.

Valid Usage

- <a href="#VUID-VkDescriptorBufferInfo-offset-00340"
  id="VUID-VkDescriptorBufferInfo-offset-00340"></a>
  VUID-VkDescriptorBufferInfo-offset-00340  
  `offset` **must** be less than the size of `buffer`

- <a href="#VUID-VkDescriptorBufferInfo-range-00341"
  id="VUID-VkDescriptorBufferInfo-range-00341"></a>
  VUID-VkDescriptorBufferInfo-range-00341  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be
  greater than `0`

- <a href="#VUID-VkDescriptorBufferInfo-range-00342"
  id="VUID-VkDescriptorBufferInfo-range-00342"></a>
  VUID-VkDescriptorBufferInfo-range-00342  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be less
  than or equal to the size of `buffer` minus `offset`

- <a href="#VUID-VkDescriptorBufferInfo-buffer-02998"
  id="VUID-VkDescriptorBufferInfo-buffer-02998"></a>
  VUID-VkDescriptorBufferInfo-buffer-02998  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, `buffer` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDescriptorBufferInfo-buffer-02999"
  id="VUID-VkDescriptorBufferInfo-buffer-02999"></a>
  VUID-VkDescriptorBufferInfo-buffer-02999  
  If `buffer` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `offset`
  **must** be zero and `range` **must** be `VK_WHOLE_SIZE`

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorBufferInfo-buffer-parameter"
  id="VUID-VkDescriptorBufferInfo-buffer-parameter"></a>
  VUID-VkDescriptorBufferInfo-buffer-parameter  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `buffer`
  **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorBufferInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
