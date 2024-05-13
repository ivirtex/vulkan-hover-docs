# VK_SEC_amigo_profiling(3) Manual Page

## Name

VK_SEC_amigo_profiling - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

486

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Ralph Potter r_potter

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-07-29

**IP Status**  
No known IP claims.

**Contributors**  
- Ralph Potter, Samsung

- Sangrak Oh, Samsung

- Jinku Kang, Samsung

## <a href="#_description" class="anchor"></a>Description

This extension is intended to communicate information from layered API
implementations such as ANGLE to internal proprietary system schedulers.
It has no behavioral implications beyond enabling more intelligent
behavior from the system scheduler.

Application developers should avoid using this extension. It is
documented solely for the benefit of tools and layer developers, who may
need to manipulate `pNext` chains that include these structures.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>There is currently no specification language written for this
extension. The links to APIs defined by the extension are to stubs that
only include generated content such as API declarations and implicit
valid usage statements.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This extension is only intended for use in specific embedded
environments with known implementation details, and is therefore
undocumented.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceAmigoProfilingFeaturesSEC](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAmigoProfilingFeaturesSEC.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html):

  - [VkAmigoProfilingSubmitInfoSEC](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAmigoProfilingSubmitInfoSEC.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_SEC_AMIGO_PROFILING_EXTENSION_NAME`

- `VK_SEC_AMIGO_PROFILING_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_AMIGO_PROFILING_SUBMIT_INFO_SEC`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_AMIGO_PROFILING_FEATURES_SEC`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-07-29 (Ralph Potter)

  - Initial specification

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_SEC_amigo_profiling"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
