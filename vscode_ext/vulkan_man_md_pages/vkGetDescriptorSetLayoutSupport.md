# vkGetDescriptorSetLayoutSupport(3) Manual Page

## Name

vkGetDescriptorSetLayoutSupport - Query whether a descriptor set layout
can be created



## <a href="#_c_specification" class="anchor"></a>C Specification

To query information about whether a descriptor set layout **can** be
created, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetDescriptorSetLayoutSupport(
    VkDevice                                    device,
    const VkDescriptorSetLayoutCreateInfo*      pCreateInfo,
    VkDescriptorSetLayoutSupport*               pSupport);
```

or the equivalent command

``` c
// Provided by VK_KHR_maintenance3
void vkGetDescriptorSetLayoutSupportKHR(
    VkDevice                                    device,
    const VkDescriptorSetLayoutCreateInfo*      pCreateInfo,
    VkDescriptorSetLayoutSupport*               pSupport);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that would create the descriptor set
  layout.

- `pCreateInfo` is a pointer to a
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
  structure specifying the state of the descriptor set layout object.

- `pSupport` is a pointer to a
  [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html)
  structure, in which information about support for the descriptor set
  layout object is returned.

## <a href="#_description" class="anchor"></a>Description

Some implementations have limitations on what fits in a descriptor set
which are not easily expressible in terms of existing limits like
`maxDescriptorSet`\*, for example if all descriptor types share a
limited space in memory but each descriptor is a different size or
alignment. This command returns information about whether a descriptor
set satisfies this limit. If the descriptor set layout satisfies the
[VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html)::`maxPerSetDescriptors`
limit, this command is guaranteed to return `VK_TRUE` in
[VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html)::`supported`.
If the descriptor set layout exceeds the
[VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html)::`maxPerSetDescriptors`
limit, whether the descriptor set layout is supported is
implementation-dependent and **may** depend on whether the descriptor
sizes and alignments cause the layout to exceed an internal limit.

This command does not consider other limits such as
`maxPerStageDescriptor`\*, and so a descriptor set layout that is
supported according to this command **must** still satisfy the pipeline
layout limits such as `maxPerStageDescriptor`\* in order to be used in a
pipeline layout.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This is a <code>VkDevice</code> query rather than
<code>VkPhysicalDevice</code> because the answer <strong>may</strong>
depend on enabled features.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkGetDescriptorSetLayoutSupport-device-parameter"
  id="VUID-vkGetDescriptorSetLayoutSupport-device-parameter"></a>
  VUID-vkGetDescriptorSetLayoutSupport-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDescriptorSetLayoutSupport-pCreateInfo-parameter"
  id="VUID-vkGetDescriptorSetLayoutSupport-pCreateInfo-parameter"></a>
  VUID-vkGetDescriptorSetLayoutSupport-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
  structure

- <a href="#VUID-vkGetDescriptorSetLayoutSupport-pSupport-parameter"
  id="VUID-vkGetDescriptorSetLayoutSupport-pSupport-parameter"></a>
  VUID-vkGetDescriptorSetLayoutSupport-pSupport-parameter  
  `pSupport` **must** be a valid pointer to a
  [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html),
[VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDescriptorSetLayoutSupport"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
