# VK_EXT_robustness2(3) Manual Page

## Name

VK_EXT_robustness2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

287

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Liam Middlebrook <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_robustness2%5D%20@liam-middlebrook%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_robustness2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>liam-middlebrook</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-01-29

**IP Status**  
No known IP claims.

**Contributors**  
- Liam Middlebrook, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds stricter requirements for how out of bounds reads
and writes are handled. Most accesses **must** be tightly
bounds-checked, out of bounds writes **must** be discarded, out of bound
reads **must** return zero. Rather than allowing multiple possible
(0,0,0,x) vectors, the out of bounds values are treated as zero, and
then missing components are inserted based on the format as described in
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-conversion-to-rgba"
target="_blank" rel="noopener">Conversion to RGBA</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input-extraction"
target="_blank" rel="noopener">vertex input attribute extraction</a>.

These additional requirements **may** be expensive on some
implementations, and should only be enabled when truly necessary.

This extension also adds support for “null descriptors”, where
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **can** be used instead of a valid
handle. Accesses to null descriptors have well-defined behavior, and do
not rely on robustness.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRobustness2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRobustness2FeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceRobustness2PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRobustness2PropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_ROBUSTNESS_2_EXTENSION_NAME`

- `VK_EXT_ROBUSTNESS_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_PROPERTIES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1.  Why do
    [VkPhysicalDeviceRobustness2PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRobustness2PropertiesEXT.html)::`robustUniformBufferAccessSizeAlignment`
    and
    [VkPhysicalDeviceRobustness2PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRobustness2PropertiesEXT.html)::`robustStorageBufferAccessSizeAlignment`
    exist?

**RESOLVED**: Some implementations cannot efficiently tightly
bounds-check all buffer accesses. Rather, the size of the bound range is
padded to some power of two multiple, up to 256 bytes for uniform
buffers and up to 4 bytes for storage buffers, and that padded size is
bounds-checked. This is sufficient to implement D3D-like behavior,
because D3D only allows binding whole uniform buffers or ranges that are
a multiple of 256 bytes, and D3D raw and structured buffers only support
32-bit accesses.

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-11-01 (Jeff Bolz, Liam Middlebrook)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_robustness2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
